// Package lsp provides the core logic for the "lsp" CLI command.
package lsp

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Run() error {
	// get current directory
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	maxDigits, paths, err := listFiles(wd)
	if err != nil {
		return err
	}

	if err := printLines(maxDigits, paths, os.Stdout); err != nil {
		return err
	}

	return nil
}

func listFiles(dir string) (int, []string, error) {
	// get files and directories
	entries, err := os.ReadDir(dir)
	if err != nil {
		return 0, nil, err
	}

	var files []string
	var maxDigits int
	for _, e := range entries {
		if e.IsDir() {
			continue
		}

		i, err := e.Info()
		if err != nil {
			return 0, nil, err
		}

		// get maximum digits in the file
		maxDigits = max(maxDigits, digits(i.Size()))
		files = append(files, filepath.Join(dir, e.Name()))
	}

	return maxDigits, files, nil
}

func printLines(maxDigits int, paths []string, w io.Writer) error {
	for _, p := range paths {
		i, err := os.Stat(p)
		if err != nil {
			return err
		}

		size := i.Size()
		// align to maximum digits
		space := strings.Repeat(" ", maxDigits-digits(size))
		fmt.Fprintln(w, size, space, p)
	}

	return nil
}

func digits(n int64) int {
	// calc argument digits
	s := fmt.Sprintf("%d", n)
	return len(s)
}
