package ruleguard

import (
	"go/ast"
	"go/types"
	"regexp"

	"github.com/quasilyte/go-ruleguard/ruleguard/quasigo"
	"github.com/quasilyte/go-ruleguard/ruleguard/typematch"
	"github.com/quasilyte/gogrep"
	"github.com/quasilyte/gogrep/nodetag"
)

type goRuleSet struct {
	universal *scopedGoRuleSet

	groups map[string]*GoRuleGroup // To handle redefinitions
}

type scopedGoRuleSet struct {
	categorizedNum int
	rulesByTag     [nodetag.NumBuckets][]goRule
	commentRules   []goCommentRule
}

type goCommentRule struct {
	base          goRule
	pat           *regexp.Regexp
	captureGroups bool
}

type goRule struct {
	group      *GoRuleGroup
	line       int
	pat        *gogrep.Pattern
	msg        string
	location   string
	suggestion string
	filter     matchFilter
	do         *quasigo.Func
}

type matchFilterResult string

func (s matchFilterResult) Matched() bool { _ = "STUB: not implemented"; return false }

func (s matchFilterResult) RejectReason() string { _ = "STUB: not implemented"; return "" }

type filterFunc func(*filterParams) matchFilterResult

type matchFilter struct {
	src string
	fn  func(*filterParams) matchFilterResult
}

type filterParams struct {
	ctx      *RunContext
	filename string
	imports  map[string]struct{}
	env      *quasigo.EvalEnv

	importer       *goImporter
	gogrepSubState *gogrep.MatcherState
	typematchState *typematch.MatcherState

	match    matchData
	nodePath *nodePath

	nodeText   func(n ast.Node) []byte
	nodeString func(n ast.Node) string

	deadcode bool

	currentFunc *ast.FuncDecl

	// varname is set only for custom filters before bytecode function is called.
	varname string

	// Both of these are Do() function related fields.
	reportString  string
	suggestString string
}

func (params *filterParams) subNode(name string) ast.Node {
	_ = "STUB: not implemented"
	return *new(ast.Node)
}

func (params *filterParams) subExpr(name string) ast.Expr {
	_ = "STUB: not implemented"
	return *new(ast.Expr)
}

func (params *filterParams) typeofNode(n ast.Node) types.Type {
	_ = "STUB: not implemented"
	return *new(types.Type)
}

func mergeRuleSets(toMerge []*goRuleSet) (*goRuleSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func appendScopedRuleSet(dst, src *scopedGoRuleSet) *scopedGoRuleSet {
	_ = "STUB: not implemented"
	return nil
}

func cloneRuleSlice(slice []goRule) []goRule { _ = "STUB: not implemented"; return nil }
