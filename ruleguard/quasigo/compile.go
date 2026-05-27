package quasigo

import (
	"go/ast"
	"go/constant"
	"go/types"
)

var voidType = &types.Tuple{}

func compile(ctx *CompileContext, fn *ast.FuncDecl) (compiled *Func, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// not our panic

func compileFunc(ctx *CompileContext, fn *ast.FuncDecl) *Func {
	_ = "STUB: not implemented"
	return nil
}

type compiler struct {
	ctx *CompileContext

	fnType  *types.Signature
	retType types.Type

	lastOp opcode

	locals           map[string]int
	constantsPool    map[interface{}]int
	intConstantsPool map[int]int

	params    map[string]int
	intParams map[string]int

	code         []byte
	constants    []interface{}
	intConstants []int

	breakTarget    *label
	continueTarget *label

	labels []*label
}

type label struct {
	targetPos int
	sources   []int
}

type compileError string

func (e compileError) Error() string { _ = "STUB: not implemented"; return "" }

func (cl *compiler) compileFunc(fn *ast.FuncDecl) *Func { _ = "STUB: not implemented"; return nil }

func (cl *compiler) compileStmt(stmt ast.Stmt) { _ = "STUB: not implemented"; return }

func (cl *compiler) compileIncDecStmt(stmt *ast.IncDecStmt) { _ = "STUB: not implemented"; return }

func (cl *compiler) compileBranchStmt(branch *ast.BranchStmt) { _ = "STUB: not implemented"; return }

func (cl *compiler) compileExprStmt(stmt *ast.ExprStmt) { _ = "STUB: not implemented"; return }

func (cl *compiler) compileForStmt(stmt *ast.ForStmt) { _ = "STUB: not implemented"; return }

// Will be implemented later; probably when the max number of locals will be lifted.

// `for <cond> { ... }`

// `for { ... }`

func (cl *compiler) compileIfStmt(stmt *ast.IfStmt) { _ = "STUB: not implemented"; return }

func (cl *compiler) compileAssignStmt(assign *ast.AssignStmt) { _ = "STUB: not implemented"; return }

func (cl *compiler) isParamName(varname string) bool { _ = "STUB: not implemented"; return false }

func (cl *compiler) getLocal(v ast.Expr, varname string) int { _ = "STUB: not implemented"; return 0 }

func (cl *compiler) compileReturnStmt(ret *ast.ReturnStmt) { _ = "STUB: not implemented"; return }

func (cl *compiler) compileExpr(e ast.Expr) { _ = "STUB: not implemented"; return }

func (cl *compiler) compileSelectorExpr(e *ast.SelectorExpr) { _ = "STUB: not implemented"; return }

func (cl *compiler) compileBinaryExpr(e *ast.BinaryExpr) { _ = "STUB: not implemented"; return }

func (cl *compiler) compileIntBinaryOp(e *ast.BinaryExpr, op opcode, typ types.Type) {
	_ = "STUB: not implemented"
	return
}

func (cl *compiler) compileSliceExpr(slice *ast.SliceExpr) { _ = "STUB: not implemented"; return }

// No need to do slicing, its no-op `s[:]`.

func (cl *compiler) compileBuiltinCall(fn *ast.Ident, call *ast.CallExpr) {
	_ = "STUB: not implemented"
	return
}

func (cl *compiler) compileCallExpr(call *ast.CallExpr) { _ = "STUB: not implemented"; return }

// TODO: just use Func.FullName as a key?

func (cl *compiler) compileCall(key funcKey, sig *types.Signature, args []ast.Expr) bool {
	_ = "STUB: not implemented"
	return false
}

func (cl *compiler) compileNativeCall(key funcKey, variadic int, funcExpr ast.Expr, args []ast.Expr) bool {
	_ = "STUB: not implemented"
	return false
}

// Check that it's not a f(g()) call, where g() returns
// a multi-value result; we can't compile that yet.

// int-typed values should appear in the interface{}-typed
// objects slice, so we get all variadic args placed in one place.

// Even if len(variadicArgs) is 0, we still need to overwrite
// the old variadicLen value, so the variadic func is not confused
// by some unrelated value.

func (cl *compiler) compileUnaryOp(op opcode, e *ast.UnaryExpr) { _ = "STUB: not implemented"; return }

func (cl *compiler) compileBinaryOp(op opcode, e *ast.BinaryExpr) {
	_ = "STUB: not implemented"
	return
}

func (cl *compiler) compileOr(e *ast.BinaryExpr) { _ = "STUB: not implemented"; return }

func (cl *compiler) compileAnd(e *ast.BinaryExpr) { _ = "STUB: not implemented"; return }

func (cl *compiler) compileIdent(ident *ast.Ident) { _ = "STUB: not implemented"; return }

func (cl *compiler) compileConstantValue(source ast.Expr, cv constant.Value) {
	_ = "STUB: not implemented"
	return
}

func (cl *compiler) internIntConstant(v int) int { _ = "STUB: not implemented"; return 0 }

func (cl *compiler) internConstant(v interface{}) int { _ = "STUB: not implemented"; return 0 }

func (cl *compiler) linkJumps() { _ = "STUB: not implemented"; return }

func (cl *compiler) newLabel() *label { _ = "STUB: not implemented"; return nil }

func (cl *compiler) bindLabel(l *label) { _ = "STUB: not implemented"; return }

func (cl *compiler) emit(op opcode) { _ = "STUB: not implemented"; return }

func (cl *compiler) emitJump(op opcode, l *label) { _ = "STUB: not implemented"; return }

func (cl *compiler) emit8(op opcode, arg8 int) { _ = "STUB: not implemented"; return }

func (cl *compiler) emit16(op opcode, arg16 int) { _ = "STUB: not implemented"; return }

func (cl *compiler) errorUnsupportedType(e ast.Node, typ types.Type, where string) compileError {
	_ = "STUB: not implemented"
	return *new(compileError)
}

func (cl *compiler) errorf(n ast.Node, format string, args ...interface{}) compileError {
	_ = "STUB: not implemented"
	return *new(compileError)
}

func (cl *compiler) isUncondJump(op opcode) bool { _ = "STUB: not implemented"; return false }

func (cl *compiler) isSupportedType(typ types.Type) bool { _ = "STUB: not implemented"; return false }

// 1. Pointers to structs are supported.

// 2. Some of the basic types are supported.
// TODO: support byte/uint8 and maybe float64.

// 3. Interfaces are supported.
