// Package cl provides the core logic for the "cl" CLI command.
package cl

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func Run() error {
	if err := validate(); err != nil {
		return err
	}

	// read input via pipe
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "read error:", err)
		os.Exit(1)
	}
	dataStr := strings.TrimRight(string(data), "\r\n")

	cmd := exec.Command("xsel", "--clipboard")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	// write stdin of xsel
	_, _ = stdin.Write([]byte(dataStr))
	stdin.Close()

	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}

func validate() error {
	if err := requireXselCommand(); err != nil {
		return err
	}

	if err := requireIsPiped(); err != nil {
		return err
	}
	return nil
}

func requireXselCommand() error {
	if _, err := exec.LookPath("xsel"); err != nil {
		return fmt.Errorf("xsel command not found: \nplease install (sudo apt install xsel)")
	}
	return nil
}

func requireIsPiped() error {
	stat, _ := os.Stdin.Stat()
	isPiped := (stat.Mode() & os.ModeCharDevice) == 0

	if !isPiped {
		return fmt.Errorf("no pipe input detected: \nusage: echo \"text\" | cl")
	}
	return nil
}
