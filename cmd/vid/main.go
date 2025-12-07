package main

import (
	"fmt"
	"os"

	"github.com/kagami-tsukimura/go-armory/internal/vid"
)

func main() {
	if err := vid.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
