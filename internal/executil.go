// Package internal provides helper functions "executil" commands.
package internal

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func RunCmd(name string, args ...string) error {
	return RunCmdStd(name, args, os.Stdout)
}

func RunCmdStd(name string, args []string, stdout io.Writer) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("command %s failed: %w", name, err)
	}

	return nil
}
