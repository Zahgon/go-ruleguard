package ruleguard

import (
	"go/ast"
)

type nodePath struct {
	stack []ast.Node
}

func newNodePath() *nodePath { _ = "STUB: not implemented"; return nil }

func (p nodePath) String() string { _ = "STUB: not implemented"; return "" }

func (p *nodePath) Parent() ast.Node { _ = "STUB: not implemented"; return *new(ast.Node) }

func (p *nodePath) Current() ast.Node { _ = "STUB: not implemented"; return *new(ast.Node) }

func (p *nodePath) NthParent(n int) ast.Node { _ = "STUB: not implemented"; return *new(ast.Node) }

func (p *nodePath) Len() int { _ = "STUB: not implemented"; return 0 }

func (p *nodePath) Push(n ast.Node) { _ = "STUB: not implemented"; return }

func (p *nodePath) Pop() { _ = "STUB: not implemented"; return }
