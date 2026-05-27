package typematch

import (
	"go/ast"
	"go/types"
)

//go:generate stringer -type=patternOp
type patternOp int

const (
	opBuiltinType patternOp = iota
	opPointer
	opVar
	opVarSeq
	opSlice
	opArray
	opMap
	opChan
	opFuncNoSeq
	opFunc
	opStructNoSeq
	opStruct
	opAnyInterface
	opNamed
)

type MatcherState struct {
	typeMatches  map[string]types.Type
	int64Matches map[string]int64
}

func NewMatcherState() *MatcherState { _ = "STUB: not implemented"; return nil }

func (state *MatcherState) reset() { _ = "STUB: not implemented"; return }

type Pattern struct {
	root *pattern
}

type pattern struct {
	value interface{}
	op    patternOp
	subs  []*pattern
}

func (pat pattern) String() string { _ = "STUB: not implemented"; return "" }

type ImportsTab struct {
	imports []map[string]string
}

func NewImportsTab(initial map[string]string) *ImportsTab { _ = "STUB: not implemented"; return nil }

func (itab *ImportsTab) Lookup(pkgName string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (itab *ImportsTab) Load(pkgName, pkgPath string) { _ = "STUB: not implemented"; return }

func (itab *ImportsTab) EnterScope() { _ = "STUB: not implemented"; return }

func (itab *ImportsTab) LeaveScope() { _ = "STUB: not implemented"; return }

type Context struct {
	Itab *ImportsTab
}

const (
	varPrefix    = `ᐸvarᐳ`
	varSeqPrefix = `ᐸvar_seqᐳ`
)

func Parse(ctx *Context, s string) (*Pattern, error) { _ = "STUB: not implemented"; return nil, nil }

var (
	builtinTypeByName = map[string]types.Type{
		"bool":       types.Typ[types.Bool],
		"int":        types.Typ[types.Int],
		"int8":       types.Typ[types.Int8],
		"int16":      types.Typ[types.Int16],
		"int32":      types.Typ[types.Int32],
		"int64":      types.Typ[types.Int64],
		"uint":       types.Typ[types.Uint],
		"uint8":      types.Typ[types.Uint8],
		"uint16":     types.Typ[types.Uint16],
		"uint32":     types.Typ[types.Uint32],
		"uint64":     types.Typ[types.Uint64],
		"uintptr":    types.Typ[types.Uintptr],
		"float32":    types.Typ[types.Float32],
		"float64":    types.Typ[types.Float64],
		"complex64":  types.Typ[types.Complex64],
		"complex128": types.Typ[types.Complex128],
		"string":     types.Typ[types.String],

		"error": types.Universe.Lookup("error").Type(),

		// Aliases.
		"byte": types.Typ[types.Uint8],
		"rune": types.Typ[types.Int32],
	}

	efaceType = types.NewInterfaceType(nil, nil)
)

func parseExpr(ctx *Context, e ast.Expr) *pattern { _ = "STUB: not implemented"; return nil }

// Only unnamed seq are supported right now.

// MatchIdentical returns true if the go typ matches pattern p.
func (p *Pattern) MatchIdentical(state *MatcherState, typ types.Type) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *Pattern) matchIdenticalFielder(state *MatcherState, subs []*pattern, f fielder) bool {
	_ = "STUB: not implemented"
	// TODO: do backtracking.
	return false
}

// "Nothing left to match" stop condition.

// Lookahead for non-greedy matching.

func (p *Pattern) matchIdentical(state *MatcherState, sub *pattern, typ types.Type) bool {
	_ = "STUB: not implemented"
	return false
}

// pkg can be nil for builtin named types.
// There is no point in checking anything else as we never
// generate the opNamed for such types.

type fielder interface {
	Field(i int) *types.Var
	NumFields() int
}

type tupleFielder struct {
	x *types.Tuple
}

func (tup *tupleFielder) Field(i int) *types.Var { _ = "STUB: not implemented"; return nil }
func (tup *tupleFielder) NumFields() int         { _ = "STUB: not implemented"; return 0 }
