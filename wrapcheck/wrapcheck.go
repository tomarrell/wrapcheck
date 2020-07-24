package wrapcheck

import (
	"flag"
	"go/ast"
	"go/types"

	"github.com/davecgh/go-spew/spew"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "wrapcheck",
	Doc:  "Checks that errors returned from external packages are wrapped",
	Flags: flag.FlagSet{
		Usage: func() { panic("not implemented") },
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			ass, ok := n.(*ast.AssignStmt)
			if !ok {
				return true
			}

			if len(ass.Rhs) < 1 {
				return true
			}

			call, ok := ass.Rhs[0].(*ast.CallExpr)
			if !ok {
				return true
			}

			if !isPackageCall(call.Fun) {
				return true
			}

			spew.Dump(call)

			for _, expr := range ass.Lhs {
				typ := pass.TypesInfo.TypeOf(expr)
				if !isError(typ) {
					return true
				}

				spew.Dump(expr)
				ident, ok := expr.(*ast.Ident)
				if !ok {
					return true
				}

				spew.Dump(ident)
			}

			return true
		})
	}

	return nil, nil
}

func isPackageCall(expr ast.Expr) bool {
	_, ok := expr.(*ast.SelectorExpr)

	return ok
}

func isError(typ types.Type) bool {
	return typ.String() == "error"
}
