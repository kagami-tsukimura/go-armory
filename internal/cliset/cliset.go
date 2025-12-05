// Package cliset provides the core logic for the "cliset" CLI command.
package cliset

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kagami-tsukimura/go-armory/internal"
)

func Run() {
	home, err := validate()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := internal.RunCmd("git", "checkout", "develop"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := createCoreLogic(home); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := createCommand(home); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("create successful: %s\n", os.Args[1])
}

func validate() (string, error) {
	if err := requireArgment(); err != nil {
		return "", err
	}

	home, err := requireHomeDir()
	if err != nil {
		return "", err
	}
	return home, nil
}

func requireArgment() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("requires 1 argument\nusage: %s <filename>", os.Args[0])
	}
	return nil
}

func requireHomeDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get user home directory: %w", err)
	}
	return home, nil
}

// create core logic file
func createCoreLogic(home string) error {
	targetDir := filepath.Join(home, "Developer", "go-armory", "internal", os.Args[1])
	targetFile := os.Args[1] + ".go"
	targetFullPath := filepath.Join(targetDir, targetFile)
	fileContent := fmt.Sprintf(`// Package %s provides the core logic for the "%s" CLI command.
package %s

func Run() error {

	return nil
}
`, os.Args[1], os.Args[1], os.Args[1])

	return createFileWithDir(targetFullPath, fileContent)
}

// create command file
func createCommand(home string) error {
	targetDir := filepath.Join(home, "Developer", "go-armory", "cmd", os.Args[1])
	targetFile := "main.go"
	targetFullPath := filepath.Join(targetDir, targetFile)
	fileContent := fmt.Sprintf(`package main

import (
	"fmt"
	"os"

	"github.com/kagami-tsukimura/go-armory/internal/%s"
)

func main() {
	if err := %s.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
`, os.Args[1], os.Args[1])

	return createFileWithDir(targetFullPath, fileContent)
}

// create directory and file
func createFileWithDir(path string, content string) error {
	// directory
	dir := filepath.Dir(path)
	if err := os.Mkdir(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}
	// file
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", path, err)
	}

	return nil
}
