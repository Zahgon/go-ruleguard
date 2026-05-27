package ruleguard

import (
	"go/token"
	"go/types"

	"github.com/quasilyte/gogrep"
	"github.com/quasilyte/gogrep/nodetag"

	"github.com/quasilyte/go-ruleguard/ruleguard/ir"
	"github.com/quasilyte/go-ruleguard/ruleguard/textmatch"
	"github.com/quasilyte/go-ruleguard/ruleguard/typematch"
)

type irLoaderConfig struct {
	ctx *LoadContext

	state *engineState

	importer *goImporter

	itab *typematch.ImportsTab

	pkg *types.Package

	gogrepFset *token.FileSet

	prefix      string
	importedPkg string
}

type irLoader struct {
	state *engineState
	ctx   *LoadContext
	itab  *typematch.ImportsTab

	pkg *types.Package

	file       *ir.File
	gogrepFset *token.FileSet

	filename string
	res      *goRuleSet

	importer *goImporter

	group *GoRuleGroup

	prefix      string // For imported packages, a prefix that is added to a rule group name
	importedPkg string // Package path; only for imported packages

	imported []*goRuleSet
}

func newIRLoader(config irLoaderConfig) *irLoader { _ = "STUB: not implemented"; return nil }

func (l *irLoader) LoadFile(filename string, f *ir.File) (*goRuleSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *irLoader) importErrorf(line int, wrapped error, format string, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *irLoader) errorf(line int, wrapped error, format string, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *irLoader) loadBundle(bundle ir.BundleImport) error { _ = "STUB: not implemented"; return nil }

func (l *irLoader) loadExternFile(prefix, pkgPath, filename string) (*goRuleSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *irLoader) compileFilterFuncs(filename string, irfile *ir.File) error {
	_ = "STUB: not implemented"
	return nil
}

// If this ever happens, user will get unexpected error
// lines for it; but we should trust that 99.9% errors
// should be caught at irconv phase so we get a valid Go
// source here as well?

func (l *irLoader) loadRuleGroup(group *ir.RuleGroup) error { _ = "STUB: not implemented"; return nil }

// Skip this group

// Should never happen

func (l *irLoader) loadRule(group *ir.RuleGroup, rule *ir.Rule) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *irLoader) loadCommentRule(resultProto goRule, rule *ir.Rule, src string, line int) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *irLoader) gogrepCompile(group *ir.RuleGroup, src string) (*gogrep.Pattern, gogrep.PatternInfo, error) {
	_ = "STUB: not implemented"
	return nil, *new(gogrep.PatternInfo), nil
}

func (l *irLoader) loadSyntaxRule(group *ir.RuleGroup, resultProto goRule, filterInfo filterInfo, rule *ir.Rule, src string, line int) error {
	_ = "STUB: not implemented"
	return nil
}

// OK: a predefined var for the "entire match"

func (l *irLoader) unwrapTypeExpr(filter ir.FilterExpr) (types.Type, error) {
	_ = "STUB: not implemented"
	return *new(types.Type), nil
}

func (l *irLoader) unwrapFuncRefExpr(filter ir.FilterExpr) (*types.Func, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: implement this.

func (l *irLoader) unwrapInterfaceExpr(filter ir.FilterExpr) (*types.Interface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *irLoader) unwrapRegexpExpr(filter ir.FilterExpr) (textmatch.Pattern, error) {
	_ = "STUB: not implemented"
	return *new(textmatch.Pattern), nil
}

func (l *irLoader) unwrapNodeTagExpr(filter ir.FilterExpr) (nodetag.Value, error) {
	_ = "STUB: not implemented"
	return *new(nodetag.Value), nil
}

func (l *irLoader) unwrapStringExpr(filter ir.FilterExpr) string {
	_ = "STUB: not implemented"
	return ""
}

func (l *irLoader) stringToBasicKind(s string) types.BasicInfo {
	_ = "STUB: not implemented"
	return *new(types.BasicInfo)
}

func (l *irLoader) newFilter(filter ir.FilterExpr, info *filterInfo) (matchFilter, error) {
	_ = "STUB: not implemented"
	return *new(matchFilter), nil
}

// OK.

func (l *irLoader) newBinaryExprFilter(filter ir.FilterExpr, info *filterInfo) (matchFilter, error) {
	_ = "STUB: not implemented"
	return *new(matchFilter), nil
}

// If constexpr is on the LHS, move it to the right, so the code below
// can imply constants being on the RHS all the time.

// Just a precaution: if we ever have a float values here,
// we may not want to rearrange anything.

// Simple commutative ops. Just swap the args.

type filterInfo struct {
	Vars map[string]struct{}

	group *ir.RuleGroup
}
