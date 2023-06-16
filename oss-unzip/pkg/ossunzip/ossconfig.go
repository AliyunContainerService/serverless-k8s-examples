package ossunzip

import (
	"fmt"
	"os"

	"k8s.io/klog"
)

const (
	ossConfigDefaultPath  = "/app/config/oss.config"
	ossDefaultCommandPath = "/app/bin/ossutil"
)

var (
	ossConfigPath  = ossConfigDefaultPath
	ossCommandPath = ossDefaultCommandPath
)

var ossConfigTemplate = `
[Credentials]
language=en
accessKeyID=%s
accessKeySecret=%s
endpoint=%s
`

func initOssConfig() error {
	ak := os.Getenv("OSS_ACCESSKEYID")
	if ak == "" {
		klog.Error("OSS_ACCESSKEYID is empty")
		return fmt.Errorf("OSS_ACCESSKEYID is empty")
	}

	sk := os.Getenv("OSS_ACCESSKEYSECRET")
	if sk == "" {
		klog.Error("OSS_ACCESSKEYSECRET is empty")
		return fmt.Errorf("OSS_ACCESSKEYSECRET is empty")
	}

	endpoint := os.Getenv("OSS_ENDPOINT")
	if endpoint == "" {
		klog.Error("OSS_ENDPOINT is empty")
		return fmt.Errorf("OSS_ENDPOINT is empty")
	}

	if v := os.Getenv("OSS_CONFIG_PATH"); v != "" {
		ossConfigPath = v
	}
	if v := os.Getenv("OSS_COMMAND_PATH"); v != "" {
		ossCommandPath = v
	}
	ossConfig := fmt.Sprintf(ossConfigTemplate, ak, sk, endpoint)
	if err := os.WriteFile(ossConfigPath, []byte(ossConfig), 0644); err != nil {
		klog.Errorf("write oss config error: %v", err)
		return err
	}

	return nil
}
