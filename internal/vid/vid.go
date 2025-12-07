// Package vid provides the core logic for the "vid" CLI command.
package vid

import (
	"fmt"
	"os"

	"github.com/kagami-tsukimura/go-armory/internal"
)

func Run() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("requires 1 argument\nusage: %s <filename>", os.Args[0])
	}

	if err := internal.RunCmd("xdg-open", os.Args[1]); err != nil {
		return err
	}

	return nil
}
