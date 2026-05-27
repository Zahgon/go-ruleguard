//go:build ignore
// +build ignore

package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

func implementsStringer(ctx *dsl.VarFilterContext) bool { _ = "STUB: not implemented"; return false }

func stringerLiteral(m dsl.Matcher) { _ = "STUB: not implemented"; return }
