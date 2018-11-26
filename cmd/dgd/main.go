package main

import (
	"github.com/natumn/dgd"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(dgd.Analyzer)
}
