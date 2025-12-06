package main

import (
	"fmt"
	"os"

	"github.com/kagami-tsukimura/go-armory/internal/shut"
)

func main() {
	if err := shut.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
