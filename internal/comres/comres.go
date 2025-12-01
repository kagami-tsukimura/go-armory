// Package comres provides the core logic for the "comres" CLI command.
package comres

import "github.com/kagami-tsukimura/go-armory/internal"

func Run() {
	internal.RunCmd("git", "reset", "--soft", "HEAD^")
}
