package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

//doc:summary suggests sorting function alternatives
//doc:before  sort.Slice(xs, func(i, j int) bool { return xs[i] < xs[j] })
//doc:after   sort.Ints(xs)
//doc:tags    refactor
func sortFuncs(m dsl.Matcher) { _ = "STUB: not implemented"; return }
