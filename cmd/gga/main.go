package main

import (
	"fmt"
	"os"

	"github.com/kagami-tsukimura/go-armory/internal/gga"
)

func main() {
	if err := gga.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
