// Package shut provides the core logic for the "shut" CLI command.
package shut

import "github.com/kagami-tsukimura/go-armory/internal"

func Run() error {
	if err := internal.RunCmd("pkill", "code"); err != nil {
		return err
	}

	if err := internal.RunCmd("sudo", "shutdown", "-h", "now"); err != nil {
		return err
	}

	return nil
}
