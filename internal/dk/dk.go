// Package dk provides the core logic for the "dk" CLI command.
package dk

import (
	"fmt"
	"os"

	"github.com/kagami-tsukimura/go-armory/internal"
)

func Run() {
	if err := internal.RunCmd("docker", "compose", "ps"); err != nil {
		if err := internal.RunCmd("docker", "ps"); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
