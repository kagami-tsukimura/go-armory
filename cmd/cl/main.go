package main

import (
	"fmt"
	"os"

	"github.com/kagami-tsukimura/go-armory/internal/cl"
)

func main() {
	if err := cl.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
