package wrapcheck

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

const errPkgName = "errors"

var Analyzer = &analysis.Analyzer{
	Name: "wrapcheck",
	Doc:  "Checks that errors returned from external packages are wrapped",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		// TODO store the last error as
		var _ *ast.AssignStmt

		ast.Inspect(file, func(n ast.Node) bool {
			if _, ok := n.(*ast.AssignStmt); ok {
				// TODO save the most recent error assignment
				return true
			}

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

				var (
					call *ast.CallExpr
				)

				// Try to pull out an *ast.CallExpr from either a short assignment `:=`
				// or a long assignment `var`/`const`
				if ass, ok := ident.Obj.Decl.(*ast.AssignStmt); ok {
					// first check for a short assignment
					call, ok = ass.Rhs[0].(*ast.CallExpr)
					if !ok {
						return true
					}
				} else if vSpec, ok := ident.Obj.Decl.(*ast.ValueSpec); ok {
					// check for a long assignment or const
					if len(vSpec.Values) < 1 {
						return true
					}
					call, ok = vSpec.Values[0].(*ast.CallExpr)
					if !ok {
						return true
					}
				} else {
					return true
				}

				sel, ok := call.Fun.(*ast.SelectorExpr)
				if !ok {
					return true
				}

				// The package of the function that we are calling which returns the
				// error
				funcPkg := pass.TypesInfo.ObjectOf(sel.Sel).Pkg()

				// If it's not a package name, then we should check the selector to
				// make sure that it's an identifier from the same package
				if pass.Pkg.Path() == funcPkg.Path() {
					return true
				} else if funcPkg.Name() == errPkgName {
					// Ignore the error if it's returned by something in the "errors"
					// package, e.g. errors.New(...)
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
