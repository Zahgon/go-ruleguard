package goutil

import (
	"go/ast"
	"go/types"
)

func ResolveFunc(info *types.Info, callable ast.Expr) (ast.Expr, *types.Func) {
	_ = "STUB: not implemented"
	return *new(ast.Expr), nil
}
