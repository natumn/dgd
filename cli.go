package main

import (
	"fmt"
	"os"
)

type CLI struct {
	Version string
	Name    string
}

func New() CLI {
	return CLI{
		Version: "0.1.0",
		Name:    "ddog",
	}
}

type Command interface {
	Run(args []string) int
	// Exit(status int) int
}

func (c *CLI) Run() int {
	if len(os.Args) < 2 {
		return exit(1)
	}

	subCmd := os.Args[1]
	args := os.Args[2:]

	var cmd Command
	switch subCmd {
	case "help":
		cmd = NewHelpCmd(c.Name)
	case "create":
		cmd = NewDocCmd()
	case "show":
		cmd = NewViewerCmd()
	case "version":
		cmd = NewVersionCmd(c.Name, c.Version)
	default:
		// exceptinal command, return standard output error.
		return exit(1)
	}

	return cmd.Run(args)
}

func exit(status int) int {
	fmt.Printf("Failed: unknown command\n")
	return status
}
