// Package internal provides helper functions "executil" commands.
package internal

import (
	"fmt"
	"os"
	"os/exec"
)

func RunCmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("command %s failed: %w", name, err)
	}
	return nil
}
