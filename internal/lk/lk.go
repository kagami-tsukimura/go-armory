// Package lk provides the core logic for the "lk" CLI command.
package lk

import "github.com/kagami-tsukimura/go-armory/internal"

func Run() error {
	if err := internal.RunCmd("gnome-screensaver-command", "-a"); err != nil {
		return err
	}

	return nil
}
