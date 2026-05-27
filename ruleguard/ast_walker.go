package ruleguard

import (
	"go/ast"

	"github.com/quasilyte/gogrep/nodetag"
)

type astWalker struct {
	nodePath *nodePath

	filterParams *filterParams

	visit func(ast.Node, nodetag.Value)
}

func (w *astWalker) Walk(root ast.Node, visit func(ast.Node, nodetag.Value)) {
	_ = "STUB: not implemented"
	return
}

func (w *astWalker) walkIdentList(list []*ast.Ident) { _ = "STUB: not implemented"; return }

func (w *astWalker) walkExprList(list []ast.Expr) { _ = "STUB: not implemented"; return }

func (w *astWalker) walkStmtList(list []ast.Stmt) { _ = "STUB: not implemented"; return }

func (w *astWalker) walkDeclList(list []ast.Decl) { _ = "STUB: not implemented"; return }

func (w *astWalker) walk(n ast.Node) { _ = "STUB: not implemented"; return }

// TODO: handle field types.
// See #252
