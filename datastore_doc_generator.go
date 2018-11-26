package dgd

import (
	"fmt"
	"go/ast"

	"github.com/tenntenn/comment"
	"github.com/tenntenn/comment/passes/commentmap"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"golang.org/x/tools/go/ast/inspector"
)

// Doc is descption for this analysis
const Doc = "cli tool for cloud datastore entity and index documentation automatically generated from source code."

var Analyzer = &analysis.Analyzer{
	Name:     "dgd",
	Doc:      Doc,
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer, commentmap.Analyzer, structtag.Analyzer},
}

var annotation = "datastore-gen-doc: true"

func init() {
	Analyzer.Flags.StringVar(&annotation, "annotation", annotation, "annotation for explicit assignment")
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	// stags := pass.ResultOf[structtag.Analyzer].(structtag.Maps)
	cmaps := pass.ResultOf[commentmap.Analyzer].(comment.Maps)

	nodeFilter := []ast.Node{
		(*ast.GenDecl)(nil),
	}

	inspect.WithStack(nodeFilter, func(n ast.Node, push bool, stack []ast.Node) bool {
		if !push {
			return false
		}

		switch n := n.(type) {
		case *ast.GenDecl:
			/*
				if inInit(pass, stack) {
					return true
				}
			*/

			if cmaps.Annotated(n, annotation) {
				fmt.Printf("Get annotation!: %+v\n", n)
				return true
			}

		}
		return false
	})

	fmt.Printf("cmaps: %+v\n", cmaps)
	return nil, nil
}
