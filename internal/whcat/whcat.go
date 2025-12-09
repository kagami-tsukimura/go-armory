// Package whcat provides the core logic for the "whcat" CLI command.
package whcat

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func Run() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("requires 1 argument\nusage: %s <filename>", os.Args[0])
	}

	programName := os.Args[1]
	filePath, err := exec.LookPath(programName)
	if err != nil {
		return err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		return err
	}

	return nil
}
