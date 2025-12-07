// Package st provides the core logic for the "st" CLI command.
package st

import "github.com/kagami-tsukimura/go-armory/internal"

func Run() error {
	if err := internal.RunCmd("git", "status"); err != nil {
		return err
	}

	return nil
}
