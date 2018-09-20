package main

import "os"

type CLI struct {
	version string
}

func main() {
	cli := NewCLI()
	os.Exit(
		cli.Run(os.Args[1:]))
}

func (c *CLI) Run(args []string) int {
	return 0
}

func NewCLI() CLI {
	return CLI{
		version: "1.0.0",
	}
}
