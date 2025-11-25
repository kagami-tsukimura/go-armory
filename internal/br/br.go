// Package br provides the core logic for the "br" CLI command.
package br

import (
	"fmt"

	"github.com/kagami-tsukimura/go-armory/internal"
)

func Run() {
	if err := internal.RunCmd("git", "branch", "-a"); err != nil {
		fmt.Println(err)
	}
}
