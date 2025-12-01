// Package dkdownup provides the core logic for the "dkdownup" CLI command.
package dkdownup

import (
	"fmt"
	"os"

	"github.com/kagami-tsukimura/go-armory/internal"
)

func Run() {
	if err := internal.RunCmd("docker", "compose", "down", "-v"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := internal.RunCmd("docker", "compose", "up", "-d"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
