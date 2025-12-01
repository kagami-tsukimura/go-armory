// Package dklog provides the core logic for the "dklog" CLI command.
package dklog

import (
	"fmt"
	"os"

	"github.com/kagami-tsukimura/go-armory/internal"
)

func Run() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "requires 1 argument\nusage: %s <service name>\n", os.Args[0])
		os.Exit(1)
	}

	if err := internal.RunCmd("docker", "compose", "logs", os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
