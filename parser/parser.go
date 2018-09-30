package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type ParsedData struct{}

func Parse(pkgs, files []string) (ParsedData, error) {
	pd := ParsedData{}
	fmt.Printf("Packages: %+v\n", pkgs)
	fmt.Printf("Files: %+v\n", files)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "./test/main.go", nil, parser.Mode(0))
	if err != nil {
		return pd, err
	}
	for _, d := range f.Decls {
		ast.Print(fset, d)
		fmt.Println()
	}

	return pd, nil
}
