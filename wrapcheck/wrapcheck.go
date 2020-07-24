package wrapcheck

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "wrapcheck",
	Doc:  "Checks that errors returned from external packages are wrapped",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			ret, ok := n.(*ast.ReturnStmt)
			if !ok {
				return true
			}

			if len(ret.Results) < 1 {
				return true
			}

			// Iterate over the values to be returned looking for errors
			for _, res := range ret.Results {
				if !isError(pass.TypesInfo.TypeOf(res)) {
					continue
				}

				ident, ok := res.(*ast.Ident)
				if !ok {
					return true
				}

				ass, ok := ident.Obj.Decl.(*ast.AssignStmt)
				if !ok {
					return true
				}

				call, ok := ass.Rhs[0].(*ast.CallExpr)
				if !ok {
					return true
				}

				sel, ok := call.Fun.(*ast.SelectorExpr)
				if !ok {
					return true
				}

				// If it's not a package name, then we should check the selector to
				// make sure that it's an identifier from the same package
				if pass.Pkg.Path() == pass.TypesInfo.ObjectOf(sel.Sel).Pkg().Path() {
					return true
				}

				if !isPackageCall(call.Fun) {
					return true
				}

				pass.Reportf(ident.NamePos, "error returned from external package is unwrapped")
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
