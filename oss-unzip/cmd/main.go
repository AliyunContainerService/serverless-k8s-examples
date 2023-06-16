package main

import (
	"context"
	"flag"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/AliyunContainerService/serverless-k8s-examples/oss-unzip/pkg/utils/signals"
	"k8s.io/klog/v2"
)

func main() {
	klog.InitFlags(nil)
	rand.Seed(time.Now().UnixNano())

	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-signals.SetupSignalHandler()
		cancel()
	}()

	cmd := newCommand(ctx)
	cmd.Flags().AddGoFlagSet(flag.CommandLine)
	flag.CommandLine.Parse([]string{})

	if err := cmd.Execute(); err != nil {
		klog.Fatal(err)
	}
}
