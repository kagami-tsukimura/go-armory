package main

import (
	"fmt"
	"os"

	"github.com/kagami-tsukimura/go-armory/internal/pipfr"
)

func main() {
	if err := pipfr.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
