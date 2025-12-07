// Package pipfr provides the core logic for the "pipfr" CLI command.
package pipfr

import (
	"os"

	"github.com/kagami-tsukimura/go-armory/internal"
)

func Run() error {
	file, err := os.Create("requirements.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	if err := internal.RunCmdStd("pip", []string{"freeze"}, file); err != nil {
		return err
	}

	return nil
}
