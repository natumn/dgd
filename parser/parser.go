package parser

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"

	"github.com/tenntenn/goq"
)

// Entity is cloud datastore structures is required for making documentation.
// TODO: entityの表現に必要なものを考える
type Entity struct {
	// Struct Struct
}

// Parse parse Go files and return cloud datastore entity structure.
func Parse(files []string) ([]*Entity, error) {
	entitys := []*Entity{}
	fs := make([]*ast.File, len(files))

	for _, file := range files {
		/*
			fi, err := os.Lstat(file)
			if err != nil {
				return nil, err
			}
		*/

		/*
			switch mode := fi.Mode(); {
			case mode.IsRegular():
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
		*/
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, file, nil, parser.Mode(0))
		if err != nil {
			return nil, err
		}

		fs = append(fs, f)
	}

	config := &types.Config{
		Importer: importer.Default(),
	}

	info := &types.Info{
		Defs: map[*ast.Ident]types.Object{},
		Uses: map[*ast.Ident]types.Object{},
	}

	if _, err := config.Check("main", fset, fs, info); err != nil {
		return nil, err
	}

	results := goq.New(fset, files, info).Query(&goq.Struct{})
	fmt.Printf("result: %+v", results)
	for i, r := range results {
		fmt.Printf("Node Pos %d: %v\n", i, r.Node.Pos())
		fmt.Printf("Object: %+v", r.Object)
	}

	return entitys, nil
}
