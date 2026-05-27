package ruleguard

import (
	"go/ast"
	"go/constant"
	"go/types"
)

var invalidType = types.Typ[types.Invalid]

func regexpHasCaptureGroups(pattern string) bool {
	_ = "STUB: not implemented"
	// regexp.Compile() uses syntax.Perl flags, so
	// we use the same flags here.
	return false
}

// true is more conservative than false

// OpCapture handles both named and unnamed capture groups.

func findDependency(pkg *types.Package, path string) *types.Package {
	_ = "STUB: not implemented"
	return nil
}

// It looks like indirect dependencies are always incomplete?
// If it's true, then we don't have to recurse here.

var typeByName = map[string]types.Type{
	// Predeclared types.
	`error`:      types.Universe.Lookup("error").Type(),
	`bool`:       types.Typ[types.Bool],
	`int`:        types.Typ[types.Int],
	`int8`:       types.Typ[types.Int8],
	`int16`:      types.Typ[types.Int16],
	`int32`:      types.Typ[types.Int32],
	`int64`:      types.Typ[types.Int64],
	`uint`:       types.Typ[types.Uint],
	`uint8`:      types.Typ[types.Uint8],
	`uint16`:     types.Typ[types.Uint16],
	`uint32`:     types.Typ[types.Uint32],
	`uint64`:     types.Typ[types.Uint64],
	`uintptr`:    types.Typ[types.Uintptr],
	`string`:     types.Typ[types.String],
	`float32`:    types.Typ[types.Float32],
	`float64`:    types.Typ[types.Float64],
	`complex64`:  types.Typ[types.Complex64],
	`complex128`: types.Typ[types.Complex128],

	// Predeclared aliases (provided for convenience).
	`byte`: types.Typ[types.Uint8],
	`rune`: types.Typ[types.Int32],
}

func typeFromString(s string) (types.Type, error) {
	_ = "STUB: not implemented"
	return *new(types.Type), nil
}

func typeFromNode(e ast.Expr) types.Type { _ = "STUB: not implemented"; return *new(types.Type) }

func intValueOf(info *types.Info, expr ast.Expr) constant.Value {
	_ = "STUB: not implemented"
	return *new(constant.Value)
}

// isPure reports whether expr is a softly safe expression and contains
// no significant side-effects. As opposed to strictly safe expressions,
// soft safe expressions permit some forms of side-effects, like
// panic possibility during indexing or nil pointer dereference.
//
// Uses types info to determine type conversion expressions that
// are the only permitted kinds of call expressions.
// Note that is does not check whether called function really
// has any side effects. The analysis is very conservative.
func isPure(info *types.Info, expr ast.Expr) bool {
	_ = "STUB: not implemented"
	// This list switch is not comprehensive and uses
	// whitelist to be on the conservative side.
	// Can be extended as needed.
	return false
}

// isPureList reports whether every expr in list is safe.
//
// See isPure.
func isPureList(info *types.Info, list []ast.Expr) bool { _ = "STUB: not implemented"; return false }

func isAddressable(info *types.Info, expr ast.Expr) bool { _ = "STUB: not implemented"; return false }

func isConstant(info *types.Info, expr ast.Expr) bool { _ = "STUB: not implemented"; return false }

func isConstantSlice(info *types.Info, expr ast.Expr) bool { _ = "STUB: not implemented"; return false }

// Matches []byte("string").

// isTypeExpr reports whether x represents a type expression.
//
// Type expression does not evaluate to any run time value,
// but rather describes a type that is used inside Go expression.
//
// For example, (*T)(v) is a CallExpr that "calls" (*T).
// (*T) is a type expression that tells Go compiler type v should be converted to.
func isTypeExpr(info *types.Info, x ast.Expr) bool { _ = "STUB: not implemented"; return false }

// Identifier may be a type expression if object
// it refers to is a type name.

func identOf(e ast.Expr) *ast.Ident { _ = "STUB: not implemented"; return nil }

func isTypeParam(typ types.Type) bool { _ = "STUB: not implemented"; return false }
