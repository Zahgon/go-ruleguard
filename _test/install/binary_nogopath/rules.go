//go:build ignore
// +build ignore

package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
	testrules "github.com/quasilyte/ruleguard-rules-test"
	subtestrules "github.com/quasilyte/ruleguard-rules-test/sub2"
)

func init() {
	dsl.ImportRules("", testrules.Bundle)
	dsl.ImportRules("", subtestrules.Bundle)
}

func isInterface(ctx *dsl.VarFilterContext) bool {
	_ = "STUB: not implemented"
	// Could be written as m["x"].Type.Underlying().Is(`interface{$*_}`) in DSL.
	return false
}

func exprUnparen(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func interfaceAddr(m dsl.Matcher) { _ = "STUB: not implemented"; return }
