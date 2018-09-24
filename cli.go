package main

import (
	"fmt"
)

var src = `package p
import _ "log"
func add(n, m int) {}
`

type CLI struct {
	version string
}

type Command interface {
	Run(args []string) int
	// Exit(status int) int
}

func New() CLI {
	return CLI{
		version: "1.0.0",
	}
}

func (c *CLI) Run(subCmd string, args []string) int {
	var cmd Command
	switch subCmd {
	case "help":
		cmd = NewDescCmd()
	case "create":
		cmd = NewDocCmd()
	case "show":
		cmd = NewViewerCmd()
	default:
		// exceptinal command, return standard output error.
		return exit(1)
	}

	return cmd.Run(args)
}

func exit(status int) int {
	fmt.Printf("failed: unknown command\n")
	return status
}
