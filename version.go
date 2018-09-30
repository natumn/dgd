package main

import "fmt"

type versionCmd struct {
	Name    string
	Version string
}

func NewVersionCmd(name, version string) Command {
	return &versionCmd{
		Name:    name,
		Version: version,
	}
}

func (c *versionCmd) Run(args []string) int {
	fmt.Printf("%v command version: %v\n", c.Name, c.Version)

	return 0
}
