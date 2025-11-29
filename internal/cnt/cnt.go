// Package cnt provides the core logic for the "cnt" CLI command.
package cnt

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func Run() {
	var targetDir string = "./"
	if len(os.Args) > 1 {
		targetDir = os.Args[1]
	}

	findCmd := exec.Command("find", targetDir, "-type", "f")
	wcCmd := exec.Command("wc", "-l")

	findStdout, err := findCmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "stdin pipe error:", err)
		os.Exit(1)
	}

	wcCmd.Stdin = findStdout

	var wcOut bytes.Buffer
	wcCmd.Stdout = &wcOut

	if err := findCmd.Start(); err != nil {
		fmt.Fprintln(os.Stderr, "find start error:", err)
		os.Exit(1)
	}

	if err := wcCmd.Start(); err != nil {
		fmt.Fprintln(os.Stderr, "wc start error:", err)
		os.Exit(1)
	}

	if err := findCmd.Wait(); err != nil {
		fmt.Fprintln(os.Stderr, "find wait error:", err)
		os.Exit(1)
	}

	if err := wcCmd.Wait(); err != nil {
		fmt.Fprintln(os.Stderr, "wc wait error:", err)
		os.Exit(1)
	}

	fmt.Printf("%s", wcOut.String())

}
