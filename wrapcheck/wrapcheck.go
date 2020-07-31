package wrapcheck

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

const errorsPkgName = "errors"

var Analyzer = &analysis.Analyzer{
	Name: "wrapcheck",
	Doc:  "Checks that errors returned from external packages are wrapped",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if _, ok := n.(*ast.AssignStmt); ok {
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
			for _, expr := range ret.Results {
				if !isError(pass.TypesInfo.TypeOf(expr)) {
					continue
				}

				ident, ok := expr.(*ast.Ident)
				if !ok {
					return true
				}

				returnPos := ident.Pos()

				// This is the original declaration
				sourceDecl := ident.Obj.Decl

				// A slice containing all the assignments which contain an identifer
				// referring to the source declaration of the error. This is to catch
				// cases where err is defined once, and then reassigned multiple times
				// within the same block. In these cases, we should check the method of
				// the most recent call.
				var assigns []*ast.AssignStmt

				// Find all assignments which have the same declaration
				ast.Inspect(file, func(n ast.Node) bool {
					if ass, ok := n.(*ast.AssignStmt); ok {
						for _, expr := range ass.Lhs {
							if !isError(pass.TypesInfo.TypeOf(expr)) {
								continue
							}
							if assIdent, ok := expr.(*ast.Ident); ok {
								if assIdent.Obj.Decl == sourceDecl {
									assigns = append(assigns, ass)
								}
							}
						}
					}

					return true
				})

				// Iterate through the assignments, comparing the token positions to
				// find the assignment that directly precedes the return position
				var mostRecentAssign *ast.AssignStmt

				for _, ass := range assigns {
					if ass.Pos() > returnPos {
						break
					}
					mostRecentAssign = ass
				}

				var (
					call *ast.CallExpr
				)

				// If the mostRecentAssign is nil, then we should check for ValueSpec
				// nodes in order to locate a possible var declaration
				if mostRecentAssign == nil {
					if vSpec, ok := ident.Obj.Decl.(*ast.ValueSpec); ok {
						// check for a long assignment or const
						if len(vSpec.Values) < 1 {
							return true
						}
						call, ok = vSpec.Values[0].(*ast.CallExpr)
						if !ok {
							return true
						}
					} else {
						// We couldn't find a short or var assign for this error return.
						// This is an error.
						panic("error: no declaration for variable")
					}
				} else {
					// Try to pull out an *ast.CallExpr from either a short assignment `:=`
					// or a long assignment `var`/`const`
					// if ass, ok := mostRecentAssign.(*ast.AssignStmt); ok {
					// first check for a short assignment
					call, ok = mostRecentAssign.Rhs[0].(*ast.CallExpr)
					if !ok {
						return true
					}
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
				} else if funcPkg.Name() == errorsPkgName {
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
	if typ == nil {
		return false
	}

	return typ.String() == "error"
}
