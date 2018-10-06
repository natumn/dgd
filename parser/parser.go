package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

// Entity is cloud datastore structures is required for making documentation.
// TODO: entityの表現に必要なものを考える
type Entity struct {
}

// Parse parse Go files and return cloud datastore entity structure.
func Parse(files []string) ([]*Entity, error) {
	entitys := []*Entity{}
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

	return entitys, nil
}
