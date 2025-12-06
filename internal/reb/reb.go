// Package reb provides the core logic for the "reb" CLI command.
package reb

import "github.com/kagami-tsukimura/go-armory/internal"

func Run() error {
	if err := internal.RunCmd("pkill", "code"); err != nil {
		return err
	}

	if err := internal.RunCmd("sudo", "shutdown", "-r", "now"); err != nil {
		return err
	}

	return nil
}
