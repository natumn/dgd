package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/natumn/ddog/parser"
)

type docCmd struct {
	name string
}

func NewDocCmd() Command {
	return &docCmd{
		name: "create",
	}
}

// NOTE: 適当に書いただけなので修正する
type Packages []string

func (p *Packages) String() string { return strings.Join(*p, ",") }

func (p *Packages) Set(pkgs string) error {
	*p = strings.Split(pkgs, ",")
	return nil
}

type Files []string

func (f *Files) String() string { return strings.Join(*f, ",") }

func (f *Files) Set(files string) error {
	*f = strings.Split(files, ",")
	return nil
}

type document struct {
	Output   string
	Type     string
	Packages Packages
	Files    Files
}

func (d *docCmd) Run(args []string) int {
	var doc document

	// TODO: make short flag(ex. package is -p)
	fs := flag.NewFlagSet(d.name, flag.ContinueOnError)
	fs.StringVar(&doc.Output, "output", "nosql.html",
		"out put file")
	fs.StringVar(&doc.Type, "type", "html",
		"file type")
	// TODO: It enable get multi value []string
	fs.Var(&doc.Packages, "package",
		"list of packages")
	fs.Var(&doc.Files, "file",
		"list of files")

	fs.Parse(args)

	// TODO: running source code parse step
	parseData, err := parser.Parse(doc.Packages, doc.Files)
	if err != nil {
		fmt.Printf("Failed parse files: %v\n", err)
		return 1
	}
	fmt.Println(parseData)

	// TODO: create document from parsed data step

	return 0
}
