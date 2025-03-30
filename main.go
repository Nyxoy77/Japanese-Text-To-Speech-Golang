package main

import (
	"os"

	"github.com/nyxoy77/go-text-to-speech/cmd"
)

func main() {
	cli := &cmd.CLI{ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
