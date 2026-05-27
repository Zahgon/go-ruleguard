package ruleguard

import (
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"

	"github.com/quasilyte/gogrep"
	"github.com/quasilyte/gogrep/nodetag"

	"github.com/quasilyte/go-ruleguard/ruleguard/quasigo"
	"github.com/quasilyte/go-ruleguard/ruleguard/textmatch"
	"github.com/quasilyte/go-ruleguard/ruleguard/typematch"
)

const filterSuccess = matchFilterResult("")

func filterFailure(reason string) matchFilterResult {
	_ = "STUB: not implemented"
	return *new(matchFilterResult)
}

func asExprSlice(x ast.Node) *gogrep.NodeSlice { _ = "STUB: not implemented"; return nil }

func exprListFilterApply(src string, list []ast.Expr, fn func(ast.Expr) bool) matchFilterResult {
	_ = "STUB: not implemented"
	return *new(matchFilterResult)
}

func makeNotFilter(src string, x matchFilter) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makeAndFilter(lhs, rhs matchFilter) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makeOrFilter(lhs, rhs matchFilter) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makeDeadcodeFilter(src string) filterFunc { _ = "STUB: not implemented"; return *new(filterFunc) }

func makeFileImportsFilter(src, pkgPath string) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makeFilePkgPathMatchesFilter(src string, re textmatch.Pattern) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makeFileNameMatchesFilter(src string, re textmatch.Pattern) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makePureFilter(src, varname string) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makeConstFilter(src, varname string) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makeConstSliceFilter(src, varname string) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makeAddressableFilter(src, varname string) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makeComparableFilter(src, varname string) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makeVarContainsFilter(src, varname string, pat *gogrep.Pattern) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makeCustomVarFilter(src, varname string, fn *quasigo.Func) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

// TODO(quasilyte): what if bytecode function panics due to the programming error?
// We should probably catch the panic here, print trace and return "false"
// from the filter (or even propagate that panic to let it crash).

func makeTypeImplementsFilter(src, varname string, iface *types.Interface) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makeTypeHasMethodFilter(src, varname string, fn *types.Func) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makeTypeHasPointersFilter(src, varname string) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makeTypeIsIntUintFilter(src, varname string, underlying bool, kind types.BasicKind) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makeTypeIsSignedFilter(src, varname string, underlying bool) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makeTypeOfKindFilter(src, varname string, underlying bool, kind types.BasicInfo) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makeTypesIdenticalFilter(src, lhsVarname, rhsVarname string) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makeRootSinkTypeIsFilter(src string, pat *typematch.Pattern) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

// TODO(quasilyte): add variadic support?

func makeTypeIsFilter(src, varname string, underlying bool, pat *typematch.Pattern) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makeTypeConvertibleToFilter(src, varname string, dstType types.Type) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makeTypeAssignableToFilter(src, varname string, dstType types.Type) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makeLineFilter(src, varname string, op token.Token, rhsVarname string) filterFunc {
	_ = "STUB: not implemented"
	// TODO(quasilyte): add variadic support.
	return *new(filterFunc)
}

func makeObjectIsVariadicParamFilter(src, varname string) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makeObjectIsGlobalFilter(src, varname string) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makeGoVersionFilter(src string, op token.Token, version GoVersion) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makeLineConstFilter(src, varname string, op token.Token, rhsValue constant.Value) filterFunc {
	_ = "STUB: not implemented"
	// TODO(quasilyte): add variadic support.
	return *new(filterFunc)
}

func makeTypeSizeConstFilter(src, varname string, op token.Token, rhsValue constant.Value) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makeTypeSizeFilter(src, varname string, op token.Token, rhsVarname string) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makeValueIntConstFilter(src, varname string, op token.Token, rhsValue constant.Value) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

// The value is unknown

func makeValueIntFilter(src, varname string, op token.Token, rhsVarname string) filterFunc {
	_ = "STUB: not implemented"
	// TODO(quasilyte): add variadic support.
	return *new(filterFunc)
}

func makeTextConstFilter(src, varname string, op token.Token, rhsValue constant.Value) filterFunc {
	_ = "STUB: not implemented"
	// TODO(quasilyte): add variadic support.
	return *new(filterFunc)
}

func makeTextFilter(src, varname string, op token.Token, rhsVarname string) filterFunc {
	_ = "STUB: not implemented"
	// TODO(quasilyte): add variadic support.
	return *new(filterFunc)
}

func makeTextMatchesFilter(src, varname string, re textmatch.Pattern) filterFunc {
	_ = "STUB: not implemented"
	// TODO(quasilyte): add variadic support.
	return *new(filterFunc)
}

func makeRootParentNodeIsFilter(src string, tag nodetag.Value) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func makeNodeIsFilter(src, varname string, tag nodetag.Value) filterFunc {
	_ = "STUB: not implemented"
	// TODO(quasilyte): add comment nodes support?
	// TODO(quasilyte): add variadic support.
	return *new(filterFunc)
}

func makeObjectIsFilter(src, varname, objectName string) filterFunc {
	_ = "STUB: not implemented"
	return *new(filterFunc)
}

func nodeIs(n ast.Node, tag nodetag.Value) bool { _ = "STUB: not implemented"; return false }

func typeHasMethod(typ types.Type, fn *types.Func) bool { _ = "STUB: not implemented"; return false }

func typeHasPointers(typ types.Type) bool { _ = "STUB: not implemented"; return false }

func findSinkRoot(params *filterParams) (ast.Node, *ast.KeyValueExpr) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

// Skip and continue.

func findContainingFunc(params *filterParams) *types.Signature {
	_ = "STUB: not implemented"
	return nil
}

func findSinkType(params *filterParams, parent ast.Node, kv *ast.KeyValueExpr, e ast.Expr) types.Type {
	_ = "STUB: not implemented"
	return *new(types.Type)
}

// TODO: some untyped int type?

// A function call argument.

// Probably a type cast.
