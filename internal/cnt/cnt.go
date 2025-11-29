// Package cnt provides the core logic for the "cnt" CLI command.
package cnt

import (
	"fmt"
	"io/fs"
	"os"
)

func Run() {
	targetDir := "./"
	// count target is os.Args[1]
	if len(os.Args) > 1 {
		targetDir = os.Args[1]
	}

	count, err := countFiles(targetDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(count)
}

func countFiles(targetDir string) (int, error) {
	var count int

	// count all files recursively
	err := fs.WalkDir(os.DirFS(targetDir), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			count++
		}
		return nil
	})

	return count, err
}
