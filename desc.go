package main

import "fmt"

type descCmd struct{}

func NewDescCmd() Command {
	return &descCmd{}
}

func (d *descCmd) Run(args []string) int {
	fmt.Println("desc!")
	return 0
}
