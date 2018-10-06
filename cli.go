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
	// TODO: ddogのコマンドのみ場合はhelpを出すか、すべてとしてプロジェクト全体のドキュメントを作成するか
	if len(os.Args) < 2 {
		fmt.Printf(usage, c.Name)
		return 0
	}

	args := os.Args[1:]
	cmd := NewCmd()

	return cmd.Run(args)
}

func exit(status int) int {
	fmt.Printf("Failed: unknown command\n")
	return status
}
