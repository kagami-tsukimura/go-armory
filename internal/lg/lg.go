// Package lg provides the core logic for the "lg" CLI command.
package lg

import (
	"fmt"

	"github.com/kagami-tsukimura/go-armory/internal"
)

func Run() {
	if err := internal.RunCmd("lazygit"); err != nil {
		fmt.Println(err)
	}
}
