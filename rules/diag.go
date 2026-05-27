package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

var Bundle = dsl.Bundle{}

//doc:summary reports always false/true conditions
//doc:before  strings.Count(s, "/") >= 0
//doc:after   strings.Count(s, "/") > 0
//doc:tags    diagnostic
func badCond(m dsl.Matcher) { _ = "STUB: not implemented"; return }
