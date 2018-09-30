package main

import "fmt"

type helpCmd struct {
	CmdName string
}

func NewHelpCmd(cmdName string) Command {
	return &helpCmd{
		CmdName: cmdName,
	}
}

const usage = `
D-dog: cli tool for cloud datastore entity and index documentation automatically generated from source code.

Usage: %s [sub command] [options]...

sub commands:
	
	create: create document from source code.
	Options: 
		--package <package name>      create document into package scope.

		--file 	  <file name>         create document form source code in this file.

		--type    <output file type>  ducument file type. ex html, json.

		--output  <file name>         ducument output file name.


	show: display document on terminal.
	Options: 
		--package <package name>      display document into package scope.

		--file    <file name>         display document form source code in this file.

		--type    <output file type>  ducument file type. ex html, json.


	help: show usage this command.

`

func (d *helpCmd) Run(args []string) int {
	fmt.Printf(usage, d.CmdName)
	return 0
}
