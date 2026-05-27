package ruleguard

import (
	"go/ast"

	"github.com/quasilyte/gogrep"
)

type matchData struct {
	match gogrep.MatchData
}

func (m matchData) Node() ast.Node { _ = "STUB: not implemented"; return *new(ast.Node) }

func (m matchData) CaptureList() []gogrep.CapturedNode { _ = "STUB: not implemented"; return nil }

func (m matchData) CapturedByName(name string) (ast.Node, bool) {
	_ = "STUB: not implemented"
	return *new(ast.Node), false
}
