package main

import (
	"fmt"
	"os"

	"github.com/kagami-tsukimura/go-armory/internal/br"
)

func main() {
	if err := br.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
