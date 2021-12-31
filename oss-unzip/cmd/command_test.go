package main

import (
	"context"
	"testing"

	"k8s.io/klog"
)

func TestUnzip(t *testing.T) {
	cmd := newCommand(context.Background())
	if err := cmd.Execute(); err != nil {
		klog.Fatal(err)
	}
}
