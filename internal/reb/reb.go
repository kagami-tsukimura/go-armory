// Package reb provides the core logic for the "reb" CLI command.
package reb

import "github.com/kagami-tsukimura/go-armory/internal"

func Run() error {
	_ = internal.RunCmd("pkill", "code")
	if err := internal.RunCmd("sudo", "shutdown", "-r", "now"); err != nil {
		return err
	}

	return nil
}
