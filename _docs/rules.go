//go:build ruleguard
// +build ruleguard

package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

// This is an example rule file for ruleguard.
//
// It's useful on its own, but its main purpose is to show you
// how one can define custom rules.
//
// In order to use it, pass this file name to a ruleguard -rule argument:
//	$ ruleguard -rules=rules.go
//
// Some rules are auto-fixable, pass the -fix argument to apply the suggested fixes:
//	$ ruleguard -fix -rules=rules.go
//
// If you want to see a "context" lines for the reported issues, use -c:
//	$ ruleguard -c=0 -rules=rules.go # Show only reported line
//	$ ruleguard -c=2 -rules=rules.go # Show reported line +2 lines of context
//
// If you want to report any issue, please do so: https://github.com/quasilyte/go-ruleguard/issues/new

func miscRules(m dsl.Matcher) {
	_ = "STUB: not implemented"
	// See http://golang.org/issue/36225
	return
}

// See https://twitter.com/dvyukov/status/1174698980208513024

// From https://github.com/dominikh/go-tools/issues/582

func exprUnparen(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func osFilepath(m dsl.Matcher) {
	_ = "STUB: not implemented"
	// path/filepath package forwards path separators so if
	// the file already uses filepath-related API it might be
	// a good idea to reduce the direct os package dependency.
	// In some cases it helps to remove the "os" package import completely.
	return
}

// See https://twitter.com/dgryski/status/1281348103505768449
func useMathBits(m dsl.Matcher) {
	_ = "STUB: not implemented"
	// RotateLeft
	return
}

// RotateRight (from docs: To rotate x right by k bits, call RotateLeftX(x, -k))

func gocriticWrapperFunc(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func gocriticNilValReturn(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func gocriticBoolExprSimplify(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func gocriticOffBy1(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func gocriticStringXBytes(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func gocriticArgOrder(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func gocriticBadCall(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func gocriticDupArg(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func gocriticDupSubExpr(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func gocriticValSwap(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func gocriticAssignOp(m dsl.Matcher) {
	_ = "STUB: not implemented"
	// We need to define ++ and -- rules before the other,
	// so they can take a precedence.
	return
}

func gocriticRegexpMust(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func gocriticMapKey(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func gocriticAppendCombine(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func gocriticYodaStyleExpr(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func gocriticUnderef(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func gocriticEmptyStringTest(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func gocriticUnslice(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func gocriticSwitchTrue(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func gocriticSloppyLen(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func gocriticNewDeref(m dsl.Matcher) {
	_ = "STUB: not implemented"
	// TODO: add missing patterns.
	return
}

func gocriticFlagDeref(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func gocriticBadLock(m dsl.Matcher) { _ = "STUB: not implemented"; return }

// `mu1` and `mu2` are added to make possible report a line where `m2` is used (with a defer)

func reviveBoolLiteralInExpr(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func gosimpleS1003(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func contextTODO(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func redundantLenCheck(m dsl.Matcher) { _ = "STUB: not implemented"; return }
