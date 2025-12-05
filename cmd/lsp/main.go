package main

import (
	"fmt"
	"os"

	"github.com/kagami-tsukimura/go-armory/internal/lsp"
)

func main() {
	if err := lsp.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
