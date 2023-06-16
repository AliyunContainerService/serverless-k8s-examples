package ossunzip

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/AliyunContainerService/serverless-k8s-examples/oss-unzip/pkg/eventbridge"
	"github.com/AliyunContainerService/serverless-k8s-examples/oss-unzip/pkg/utils/command"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
	"k8s.io/klog"
)

const (
	WebserverPort = "8080"
)

func StartWorker() error {
	klog.Info("unzip server started~~")
	return StartWebServer()
}

func StartWebServer() error {
	if err := initOssConfig(); err != nil {
		return err
	}

	r := mux.NewRouter()
	r.HandleFunc("/unzip", unzip).Methods("POST")
	r.HandleFunc("/ping/", pong)
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%s", WebserverPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	klog.Fatal(srv.ListenAndServe())
	return nil
}

func pong(w http.ResponseWriter, r *http.Request) {
	for k, vals := range r.Header {
		for _, v := range vals {
			fmt.Fprintf(w, "header %s: %s \n", k, v)
		}
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ping pong ~~\n")
}

func unzip(w http.ResponseWriter, r *http.Request) {
	for k, vals := range r.Header {
		for _, v := range vals {
			klog.Infof("request header %s: %s \n", k, v)
		}
	}

	var body []byte
	if r.Body != nil {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			klog.Errorf("An error occured reading the request body: %v", err)
			http.Error(w, "An error occured reading the request body", http.StatusInternalServerError)
			return
		}
		body = data
		klog.Infof("Request body: %s", string(body))
	} else {
		klog.Info("Request body is empty")
	}

	event := cloudevents.NewEvent()
	if err := json.Unmarshal(body, &event); err != nil {
		klog.Errorf("An error occured unmarshalling the request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, http.StatusText(http.StatusBadRequest))
		return
	}
	ossEventBts := event.Data()
	if ossEventBts == nil {
		klog.Errorf("ossEvent Body is nil")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, http.StatusText(http.StatusBadRequest))
		return
	}

	ossEvent := &eventbridge.OSSEventSpec{}
	if err := json.Unmarshal(ossEventBts, ossEvent); err != nil {
		klog.Errorf("An error occured unmarshalling the request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, http.StatusText(http.StatusBadRequest))
		return
	}

	ossBucket := ossEvent.Oss.Bucket.Name
	zipFilePath := ossEvent.Oss.Object.Key
	klog.Infof("ossEvent received, ossBucket: %s, zipFilePath: %s", ossBucket, zipFilePath)
	if err := syncFile(ossCommandPath, ossBucket, zipFilePath); err != nil {
		klog.Errorf("sync file error:%s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, http.StatusText(http.StatusInternalServerError))
		return
	}

	klog.Infof("Event: %s", event.String())
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}

func GetUUID() string {
	u2 := uuid.NewV4()
	return u2.String()
}

func syncFile(ossCmd, bucket, ossZipFileFullPathName string) error {
	// set base dir for unzip operation
	tmpUuid := GetUUID()
	zipFileName := filepath.Base(ossZipFileFullPathName)
	localPathfilePath := filepath.Join("/tmp/ossunzip/", tmpUuid)
	localZipFileFullPathName := filepath.Join(localPathfilePath, zipFileName)
	if err := os.MkdirAll(localPathfilePath, os.ModePerm); err != nil {
		return fmt.Errorf("create dir:%s error:%v", localPathfilePath, err)
	}

	defer func() {
		if err := os.RemoveAll(localPathfilePath); err != nil {
			klog.Errorf("remove dir:%s error:%v", localPathfilePath, err)
		}
	}()

	// download file
	klog.Infof("start download file:%s ", ossZipFileFullPathName)
	remotePath := fmt.Sprintf("oss://%s/%s", bucket, ossZipFileFullPathName)
	ossConfigFile := fmt.Sprintf("--config-file=%s", ossConfigPath)
	donwloadCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := command.RunCommand(ossCmd,
		[]string{"cp", remotePath, localZipFileFullPathName, ossConfigFile, "-f", "--parallel=100"}, donwloadCtx); err != nil {
		return fmt.Errorf("download file:%s error:%v", ossZipFileFullPathName, err)
	}
	klog.Infof("download:%s success", ossZipFileFullPathName)

	//  unzip file
	klog.Infof("start unzip file:%s ", ossZipFileFullPathName)
	zipFileNamePrefix := strings.TrimSuffix(zipFileName, ".zip")
	zipFileDir := fmt.Sprintf("unzip-%s", zipFileNamePrefix)
	cmd := fmt.Sprintf("cd %s && unzip -d %s %s", localPathfilePath, zipFileDir, zipFileName)
	unzipCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := command.RunCommand("sh", []string{"-c", cmd}, unzipCtx); err != nil {
		return fmt.Errorf("unzip file:%s error:%v", ossZipFileFullPathName, err)
	}
	klog.Infof("unzip file:%s success", ossZipFileFullPathName)

	// upload file
	klog.Infof("start upload file:%s ", ossZipFileFullPathName)
	unzipFilePath := filepath.Join(localPathfilePath, zipFileDir)
	files, err := ioutil.ReadDir(unzipFilePath)
	if err != nil {
		klog.Errorf("read unzipFilePath error:%v", err)
		return fmt.Errorf("read unzipFilePath error:%v", err)
	}

	// Avoid repeated path nesting, example: serverless-k8s-examples-master/serverless-k8s-examples-master/*
	if len(files) == 1 && files[0].IsDir() && files[0].Name() == zipFileNamePrefix {
		unzipFilePath = filepath.Join(unzipFilePath, zipFileNamePrefix)
	}
	unzipRemotePath := fmt.Sprintf("oss://%s/unzip/%s", bucket, zipFileNamePrefix)
	uploadCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := command.RunCommand(ossCmd, []string{"cp", unzipFilePath, unzipRemotePath, ossConfigFile, "-f", "--parallel=100", "--recursive"}, uploadCtx); err != nil {
		return fmt.Errorf("upload file:%s error:%v", ossZipFileFullPathName, err)
	}
	klog.Infof("upload unzip file:%s success", ossZipFileFullPathName)

	return nil
}
