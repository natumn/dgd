package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

type ParsedData struct{}

func Parse(files []string) (*ParsedData, error) {
	pd := ParsedData{}
	fmt.Printf("Files: %+v\n", files)

	for _, file := range files {
		fi, err := os.Lstat(file)
		if err != nil {
			return nil, err
		}

		fset := token.NewFileSet()
		switch mode := fi.Mode(); {
		case mode.IsRegular():
			f, err := parser.ParseFile(fset, file, nil, parser.Mode(0))
			if err != nil {
				return nil, err
			}

			for _, d := range f.Decls {
				ast.Print(fset, d)
				fmt.Println()
			}
		case mode.IsDir():
			pkgs, err := parser.ParseDir(fset, file, nil, 0)
			if err != nil {
				return nil, err
			}
			for _, pkg := range pkgs {
				for _, f := range pkg.Files {
					for _, d := range f.Decls {
						ast.Print(fset, d)
						fmt.Println()
					}
				}
			}
		}
	}

	return &pd, nil
}
