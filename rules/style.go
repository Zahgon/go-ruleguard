package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

//doc:summary reports redundant parentheses
//doc:before  f(x, (y))
//doc:after   f(x, y)
//doc:tags    style
func exprUnparen(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary reports empty declaration blocks
//doc:before  var ()
//doc:after   /* nothing */
//doc:tags    style
func emptyDecl(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary reports empty errors creation
//doc:before  errors.New("")
//doc:after   errors.New("can't open the cache file")
//doc:tags    style
func emptyError(m dsl.Matcher) { _ = "STUB: not implemented"; return }
