package main

import "fmt"

type docCmd struct{}

func NewDocCmd() Command {
	return &docCmd{}
}

func (d *docCmd) Run(args []string) int {
	fmt.Println("doc!")
	return 0
}
