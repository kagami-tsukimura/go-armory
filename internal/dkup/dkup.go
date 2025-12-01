// Package dkup provides the core logic for the "dkup" CLI command.
package dkup

import (
	"fmt"
	"os"

	"github.com/kagami-tsukimura/go-armory/internal"
)

func Run() {
	if err := internal.RunCmd("docker", "compose", "up", "-d"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
