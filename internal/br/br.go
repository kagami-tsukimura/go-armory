// Package br provides the core logic for the "br" CLI command.
package br

import (
	"fmt"
	"os"

	"github.com/kagami-tsukimura/go-armory/internal"
)

func Run() {
	if err := internal.RunCmd("git", "branch", "-a"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
