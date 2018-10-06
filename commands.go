package main

import (
	"flag"
	"fmt"

	"github.com/natumn/ddog/parser"
)

type Cmd struct {
	name string
}

func NewCmd() Command {
	return &Cmd{}
}

type document struct {
	Output string
	Type   string
	Files  []string
}

func (d *Cmd) Run(args []string) int {
	var doc document
	var help, version bool

	// TODO: make short flag(ex. --type is -t)
	fs := flag.NewFlagSet(d.name, flag.ContinueOnError)
	fs.StringVar(&doc.Type, "type", "html", "specify file type")
	fs.BoolVar(&help, "help", false, "display the command usage")
	fs.BoolVar(&version, "version", false, "display the command version")

	if err := fs.Parse(args); err != nil {
		fmt.Printf("Failed parse flags: %v\n", err)
		return 1
	}

	doc.Files = fs.Args()

	// TODO: running source code parse step
	_, err := parser.Parse(doc.Files)
	if err != nil {
		fmt.Printf("Failed parse files: %v\n", err)
		return 1
	}

	// TODO: create document from parsed data step

	return 0
}
