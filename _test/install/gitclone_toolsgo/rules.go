//go:build ignore
// +build ignore

package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
	testrules "github.com/quasilyte/ruleguard-rules-test"
)

func init() {
	dsl.ImportRules("testrules", testrules.Bundle)
}

func implementsStringer(ctx *dsl.VarFilterContext) bool { _ = "STUB: not implemented"; return false }

func exprUnparen(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func sprintStringer(m dsl.Matcher) { _ = "STUB: not implemented"; return }
