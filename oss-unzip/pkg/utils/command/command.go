package command

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
	"syscall"

	"k8s.io/klog/v2"
)

func RunCommand(command string, args []string, ctx context.Context) error {
	cmd := exec.CommandContext(ctx, command, args...)
	stdoutBytes := &bytes.Buffer{}
	stderrBytes := &bytes.Buffer{}
	cmd.Stdout = stdoutBytes
	cmd.Stderr = stderrBytes

	// Start command asynchronously
	err := cmd.Start()
	if err != nil {
		return err
	}

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		if exciter, ok := err.(*exec.ExitError); ok {
			// The program has exited with an exit code != 0

			// This works on both Unix and Windows. Although package
			// syscall is generally platform dependent, WaitStatus is
			// defined for both Unix and Windows and in both cases has
			// an ExitStatus() method with the same signature.
			if status, ok := exciter.Sys().(syscall.WaitStatus); ok {
				msg := fmt.Sprintf("RunCommand: %s args: %s Exit Status: %d, stdout:%s, stderr:%s", command, args, status.ExitStatus(), stdoutBytes.String(), stderrBytes.String())
				klog.Error(msg)
				return errors.New(msg)
			}
		}
		msg := fmt.Sprintf("RunCommand: %s args: %s cmd.Wait error: %v, stdout:%s stderr:%s", command, args, err.Error(), stdoutBytes.String(), stderrBytes.String())
		klog.Error(msg)
		return err
	}
	klog.Infof("RunCommand success, command: %s args: %s Result stdout:%s ,stderr:%s", command, args, stdoutBytes.String(), stderrBytes.String())
	return nil
}
