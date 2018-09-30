package main

import "fmt"

type viewerCmd struct {
	Name string
}

func NewViewerCmd() Command {
	return &viewerCmd{
		Name: "show",
	}
}

func (d *viewerCmd) Run(args []string) int {
	fmt.Println("viewer!")
	return 0
}
