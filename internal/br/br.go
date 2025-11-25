// Package br provides the core logic for the "br" CLI command.
package br

import (
	"fmt"
	"os"
	"os/exec"
)

func Run() {
	cmd := exec.Command("git", "branch", "-a")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Command Error: %v\n", err)
	}
}
