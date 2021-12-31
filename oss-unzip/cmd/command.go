package main

import (
	"context"

	unzip "github.com/AliyunContainerService/serverless-k8s-examples/oss-unzip/pkg/ossunzip"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

const (
	// cmdName is the name of the command
	cmdName = "oss-unzip"
)

func newCommand(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   cmdName,
		Short: "start workload",
		Long:  "start workload",
		RunE: func(c *cobra.Command, args []string) error {
			return runE(ctx)
		},
	}
	return cmd
}

func runE(ctx context.Context) error {
	klog.Infof("%s started ~~", cmdName)
	if err := unzip.StartWorker(); err != nil {
		klog.Errorf("start %s error:%s", cmdName, err.Error())
		return err
	}
	klog.Infof("%s stopped ~~", cmdName)
	return nil
}
