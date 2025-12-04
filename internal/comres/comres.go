// Package comres provides the core logic for the "comres" CLI command.
package comres

import (
	"fmt"
	"os"

	"github.com/kagami-tsukimura/go-armory/internal"
)

func Run() {
	if err := internal.RunCmd("git", "reset", "--soft", "HEAD^"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
