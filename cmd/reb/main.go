package main

import (
	"fmt"
	"os"

	"github.com/kagami-tsukimura/go-armory/internal/reb"
)

func main() {
	if err := reb.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
