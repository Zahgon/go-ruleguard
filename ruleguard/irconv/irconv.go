package irconv

import (
	"go/ast"
	"go/token"
	"go/types"
	"regexp"

	"github.com/quasilyte/go-ruleguard/ruleguard/ir"
)

type Context struct {
	Pkg   *types.Package
	Types *types.Info
	Fset  *token.FileSet
	Src   []byte
}

func ConvertFile(ctx *Context, f *ast.File) (result *ir.File, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// not our panic

type convError struct {
	err error
}

type localMacroFunc struct {
	name     string
	params   []string
	template ast.Expr
}

type converter struct {
	types *types.Info
	pkg   *types.Package
	fset  *token.FileSet
	src   []byte

	versionPathRe *regexp.Regexp

	group      *ir.RuleGroup
	groupFuncs []localMacroFunc

	dslPkgname string // The local name of the "ruleguard/dsl" package (usually its just "dsl")
}

func (conv *converter) errorf(n ast.Node, format string, args ...interface{}) convError {
	_ = "STUB: not implemented"
	return *new(convError)
}

func (conv *converter) ConvertFile(f *ast.File) *ir.File { _ = "STUB: not implemented"; return nil }

// Right now this list is hardcoded from the knowledge of which
// stdlib packages are supported inside the bytecode.

func (conv *converter) convertInitFunc(dst *ir.File, decl *ast.FuncDecl) {
	_ = "STUB: not implemented"
	return
}

func (conv *converter) addCustomImport(dst *ir.File, pkgPath string) {
	_ = "STUB: not implemented"
	return
}

func (conv *converter) addCustomDecl(dst *ir.File, decl ast.Decl) {
	_ = "STUB: not implemented"
	return
}

func (conv *converter) isMatcherFunc(f *ast.FuncDecl) bool { _ = "STUB: not implemented"; return false }

func (conv *converter) convertRuleGroup(decl *ast.FuncDecl) *ir.RuleGroup {
	_ = "STUB: not implemented"
	return nil
}

func (conv *converter) findLocalMacro(call *ast.CallExpr) *localMacroFunc {
	_ = "STUB: not implemented"
	return nil
}

func (conv *converter) expandMacro(macro *localMacroFunc, call *ast.CallExpr) ir.FilterExpr {
	_ = "STUB: not implemented"
	// Check that call args are OK.
	// Since "function calls" are implemented as a macro expansion here,
	// we don't allow arguments that have a non-trivial evaluation.
	return *new(ir.FilterExpr)
}

// astcopy above will copy the AST tree, but it won't update
// the associated types.Info map of const values.
// We'll try to solve that issue at least partially here.

func (conv *converter) localDefine(assign *ast.AssignStmt) { _ = "STUB: not implemented"; return }

func (conv *converter) doMatcherImport(call *ast.CallExpr) { _ = "STUB: not implemented"; return }

// Try to be at least somewhat module-aware.
// If the last path part is "/v%d", we might want to take
// the previous path part as a package name.

func (conv *converter) doMatcherImportAs(call *ast.CallExpr) { _ = "STUB: not implemented"; return }

func (conv *converter) matcherMethodName(call *ast.CallExpr) string {
	_ = "STUB: not implemented"
	return ""
}

func (conv *converter) convertDocComments(comment *ast.CommentGroup) {
	_ = "STUB: not implemented"
	return
}

// Should never happen

func (conv *converter) convertRuleExpr(call *ast.CallExpr) { _ = "STUB: not implemented"; return }

// AST patterns for Match() or regexp patterns for MatchComment().

func (conv *converter) convertFilterExpr(e ast.Expr) ir.FilterExpr {
	_ = "STUB: not implemented"
	return *new(ir.FilterExpr)
}

func (conv *converter) convertFilterExprImpl(e ast.Expr) ir.FilterExpr {
	_ = "STUB: not implemented"
	return *new(ir.FilterExpr)
}

// TODO: reuse the code with parsing At() args?

// TODO: remove this restriction.

// TODO: remove this restriction.

func (conv *converter) parseStringArg(e ast.Expr) string { _ = "STUB: not implemented"; return "" }

func (conv *converter) toStringValue(x ast.Node) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (conv *converter) inspectFilterSelector(e ast.Expr) filterExprSelector {
	_ = "STUB: not implemented"
	return *new(filterExprSelector)
}

type filterExprSelector struct {
	mapName string
	varName string
	path    string
	args    []ast.Expr
}
