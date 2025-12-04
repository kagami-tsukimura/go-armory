// Package clone provides the core logic for the "clone" CLI command.
package clone

import (
	"fmt"
	"os"

	"github.com/kagami-tsukimura/go-armory/internal"
)

func Run() {
	// os.Args[0]: command, os.Args[1]: repository, os.Args[2]: directory
	if len(os.Args) < 2 || len(os.Args) > 3 {
		err := fmt.Errorf("requires 1 or 2 arguments\nusage: %s <repository> [directory]", os.Args[0])
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// slice in parts
	cmdArgs := append([]string{"clone"}, os.Args[1:]...)
	if err := internal.RunCmd("git", cmdArgs...); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
