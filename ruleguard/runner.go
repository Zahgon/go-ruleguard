package ruleguard

import (
	"context"
	"go/ast"
	"go/build"

	"github.com/quasilyte/gogrep"
	"github.com/quasilyte/gogrep/nodetag"
)

type rulesRunner struct {
	state *engineState

	bgContext context.Context

	ctx   *RunContext
	rules *goRuleSet

	truncateLen int

	reportData ReportData

	gogrepState    gogrep.MatcherState
	gogrepSubState gogrep.MatcherState

	importer *goImporter

	filename string
	src      []byte

	// nodePath is a stack of ast.Nodes we visited to this point.
	// When we enter a new node, it's placed on the top of the stack.
	// When we leave that node, it's popped.
	// The stack is a slice that is allocated only once and reused
	// for the lifetime of the runner.
	// The only overhead it has is a slice append and pop operations
	// that are quire cheap.
	//
	// Note: we need this path to get a Node.Parent() for `$$` matches.
	// So it's used to climb up the tree there.
	// For named submatches we can't use it as the node can be located
	// deeper into the tree than the current node.
	// In those cases we need a more complicated algorithm.
	nodePath *nodePath

	filterParams filterParams
}

func newRunnerState(es *engineState) *RunnerState { _ = "STUB: not implemented"; return nil }

func (state *RunnerState) Reset() { _ = "STUB: not implemented"; return }

func newRulesRunner(ctx *RunContext, buildContext *build.Context, state *engineState, rules *goRuleSet) *rulesRunner {
	_ = "STUB: not implemented"
	return nil
}

func (rr *rulesRunner) nodeString(n ast.Node) string { _ = "STUB: not implemented"; return "" }

func (rr *rulesRunner) nodeText(n ast.Node) []byte { _ = "STUB: not implemented"; return nil }

// Go printer would panic on comments.

// Fallback to the printer.

func (rr *rulesRunner) fileBytes() []byte { _ = "STUB: not implemented"; return nil }

// TODO(quasilyte): re-use src slice?

// Assign a zero-length slice so rr.src
// is never nil during the second fileBytes call.

func (rr *rulesRunner) run(f *ast.File) error {
	_ = "STUB: not implemented"
	// If it's not empty then we're leaking memory.
	// For every Push() there should be a Pop() call.
	return nil
}

func (rr *rulesRunner) runCommentRules(comment *ast.Comment) {
	_ = "STUB: not implemented"
	// We'll need that file to create a token.Pos from the artificial offset.
	return
}

// Negative index a special case when named group captured nothing.
// Consider this pattern: `(?P<x>foo)|(bar)`.
// If we have `bar` input string, <x> will remain empty.

// Fast path: no need to save any submatches.

func (rr *rulesRunner) runRules(n ast.Node, tag nodetag.Value) {
	_ = "STUB: not implemented"
	// profiling.LabelsEnabled is constant, so labels-related
	// code should be a no-op inside normal build.
	// To enable labels, use "-tags pproflabels" build tag.
	return
}

func (rr *rulesRunner) reject(rule goRule, reason string, m matchData) {
	_ = "STUB: not implemented"
	return
}

// This rule is not being debugged

func (rr *rulesRunner) handleCommentMatch(rule goCommentRule, m matchData) bool {
	_ = "STUB: not implemented"
	return false
}

func (rr *rulesRunner) handleMatch(rule goRule, m gogrep.MatchData) bool {
	_ = "STUB: not implemented"
	return false
}

func (rr *rulesRunner) collectImports(f *ast.File) { _ = "STUB: not implemented"; return }

func (rr *rulesRunner) renderMessage(msg string, m matchData, truncate bool) string {
	_ = "STUB: not implemented"
	return ""
}

// Some captured nodes are typed, but nil.
// We can't really get their text, so skip them here.
// For example, pattern `func $_() $results { $*_ }` may
// match a nil *ast.FieldList for $results if executed
// against a function with no results.

func (rr *rulesRunner) fixedText(text []byte, n ast.Node, following string) []byte {
	_ = "STUB: not implemented"
	// pattern=`$x.y` $x=`&buf` following=`.y`
	// Insert $x as `buf`, so we get `buf.y` instead of incorrect `&buf.y`.
	return nil
}

var longTextPlaceholder = []byte("<...>")

func truncateText(s []byte, maxLen int) []byte { _ = "STUB: not implemented"; return nil }

var multiMatchTags = [nodetag.NumBuckets]bool{
	nodetag.BlockStmt:  true,
	nodetag.CaseClause: true,
	nodetag.CommClause: true,
	nodetag.File:       true,
}
