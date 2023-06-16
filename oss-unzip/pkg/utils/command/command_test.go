package command

import (
	"context"
	"testing"

	"k8s.io/klog"
)

func TestRunCommand(t *testing.T) {
	klog.InitFlags(nil)
	if err := RunCommand("bash", []string{"-c", "sleep 3s;echo 'xxx success ~~'; exit 3"}, context.Background()); err != nil {
		t.Fatalf("error:%s", err.Error())
	}
}
