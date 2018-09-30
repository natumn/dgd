package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/natumn/ddog/parser"
)

type Cmd struct {
	name string
}

func NewCmd() Command {
	return &Cmd{}
}

type Files []string

func (f *Files) String() string { return strings.Join(*f, ",") }

func (f *Files) Set(files string) error {
	*f = strings.Split(files, ",")
	return nil
}

type document struct {
	Output string
	Type   string
	Files  Files
}

func (d *Cmd) Run(args []string) int {
	var doc document

	// TODO: make short flag(ex. package is -p)
	fs := flag.NewFlagSet(d.name, flag.ContinueOnError)
	fs.StringVar(&doc.Output, "output", "nosql.html",
		"out put file")
	fs.StringVar(&doc.Type, "type", "html",
		"file type")

	fs.Parse(args)
	files := fs.Args()

	// TODO: running source code parse step
	_, err := parser.Parse(files)
	if err != nil {
		fmt.Printf("Failed parse files: %v\n", err)
		return 1
	}

	// TODO: create document from parsed data step

	return 0
}
