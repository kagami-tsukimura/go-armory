package main

import (
	"fmt"
	"os"

	"github.com/kagami-tsukimura/go-armory/internal/spd"
)

func main() {
	if err := spd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
