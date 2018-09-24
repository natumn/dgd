package main

import (
	"os"
)

func main() {
	subCmd := os.Args[1]
	args := os.Args[2:]

	cli := New()
	os.Exit(
		cli.Run(subCmd, args),
	)
}
