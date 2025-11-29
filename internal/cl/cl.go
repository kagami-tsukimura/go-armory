// Package cl provides the core logic for the "cl" CLI command.
package cl

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func Run() {
	if err := validate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var data string
	// read input via pipe
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data += scanner.Text() + "\n"
	}

	cmd := exec.Command("xsel", "--clipboard")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "stdin pipe error:", err)
		os.Exit(1)
	}

	if err := cmd.Start(); err != nil {
		fmt.Fprintln(os.Stderr, "command start error:", err)
		os.Exit(1)
	}

	// write stdin of xsel
	_, _ = stdin.Write([]byte(data))
	stdin.Close()

	if err := cmd.Wait(); err != nil {
		fmt.Fprintln(os.Stderr, "command wait error:", err)
		os.Exit(1)
	}
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
