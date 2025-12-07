// Package shut provides the core logic for the "shut" CLI command.
package shut

import "github.com/kagami-tsukimura/go-armory/internal"

func Run() error {
	_ = internal.RunCmd("pkill", "code")
	if err := internal.RunCmd("sudo", "shutdown", "-h", "now"); err != nil {
		return err
	}

	return nil
}
