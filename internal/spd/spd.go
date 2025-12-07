// Package spd provides the core logic for the "spd" CLI command.
package spd

import "github.com/kagami-tsukimura/go-armory/internal"

func Run() error {
	if err := internal.RunCmd("speedtest-cli"); err != nil {
		return err
	}

	return nil
}
