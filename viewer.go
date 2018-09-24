package main

import "fmt"

type viewerCmd struct{}

func NewViewerCmd() Command {
	return &viewerCmd{}
}

func (d *viewerCmd) Run(args []string) int {
	fmt.Println("viewer!")
	return 0
}
