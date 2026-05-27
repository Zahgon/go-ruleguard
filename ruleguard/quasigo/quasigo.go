// Package quasigo implements a Go subset compiler and interpreter.
//
// The implementation details are not part of the contract of this package.
package quasigo

import (
	"go/ast"
	"go/token"
	"go/types"
)

// TODO(quasilyte): document what is thread-safe and what not.
// TODO(quasilyte): add a readme.

// Env is used to hold both compilation and evaluation data.
type Env struct {
	// TODO(quasilyte): store both native and user func ids in one map?

	nativeFuncs        []nativeFunc
	nameToNativeFuncID map[funcKey]uint16

	userFuncs    []*Func
	nameToFuncID map[funcKey]uint16

	// debug contains all information that is only needed
	// for better debugging and compiled code introspection.
	// Right now it's always enabled, but we may allow stripping it later.
	debug *debugInfo
}

// EvalEnv is a goroutine-local handle for Env.
// To get one, use Env.GetEvalEnv() method.
type EvalEnv struct {
	nativeFuncs []nativeFunc
	userFuncs   []*Func

	Stack ValueStack
}

// NewEnv creates a new empty environment.
func NewEnv() *Env {
	_ = "STUB: not implemented"

	// GetEvalEnv creates a new goroutine-local handle of env.
	return nil
}

func (env *Env) GetEvalEnv() *EvalEnv { _ = "STUB: not implemented"; return nil }

// AddNativeMethod binds `$typeName.$methodName` symbol with f.
// A typeName should be fully qualified, like `github.com/user/pkgname.TypeName`.
// It method is defined only on pointer type, the typeName should start with `*`.
func (env *Env) AddNativeMethod(typeName, methodName string, f func(*ValueStack)) {
	_ = "STUB: not implemented"
	return
}

// AddNativeFunc binds `$pkgPath.$funcName` symbol with f.
// A pkgPath should be a full package path in which funcName is defined.
func (env *Env) AddNativeFunc(pkgPath, funcName string, f func(*ValueStack)) {
	_ = "STUB: not implemented"
	return
}

// AddFunc binds `$pkgPath.$funcName` symbol with f.
func (env *Env) AddFunc(pkgPath, funcName string, f *Func) { _ = "STUB: not implemented"; return }

// GetFunc finds previously bound function searching for the `$pkgPath.$funcName` symbol.
func (env *Env) GetFunc(pkgPath, funcName string) *Func { _ = "STUB: not implemented"; return nil }

// CompileContext is used to provide necessary data to the compiler.
type CompileContext struct {
	// Env is shared environment that should be used for all functions
	// being compiled; then it should be used to execute these functions.
	Env *Env

	Package *types.Package
	Types   *types.Info
	Fset    *token.FileSet
}

// Compile prepares an executable version of fn.
func Compile(ctx *CompileContext, fn *ast.FuncDecl) (compiled *Func, err error) {
	_ = "STUB: not implemented"
	return nil,

		// Call invokes a given function.
		// All arguments should be pushed to env.Stack prior to this call.
		//
		// Note that arguments are not popped off the stack,
		// so you can bind the args once and use Call multiple times.
		// If you want to reset arguments, do env.Stack.Reset().
		nil
}

func Call(env *EvalEnv, fn *Func) CallResult { _ = "STUB: not implemented"; return *new(CallResult) }

// CallResult is a return value of Call function.
// For most functions, Value() should be called to get the actual result.
// For int-typed functions, IntValue() should be used instead.
type CallResult struct {
	value       interface{}
	scalarValue uint64
}

// Value unboxes an actual call return value.
// For int results, use IntValue().
func (res CallResult) Value() interface{} {
	_ = "STUB: not implemented"

	// IntValue unboxes an actual call return value.
	return nil
}

func (res CallResult) IntValue() int { _ = "STUB: not implemented"; return 0 }

// Disasm returns the compiled function disassembly text.
// This output is not guaranteed to be stable between versions
// and should be used only for debugging purposes.
func Disasm(env *Env, fn *Func) string { _ = "STUB: not implemented"; return "" }

// Func is a compiled function that is ready to be executed.
type Func struct {
	code []byte

	constants    []interface{}
	intConstants []int

	numObjectParams int
	numIntParams    int

	name string
}

// ValueStack is used to manipulate runtime values during the evaluation.
// Function arguments are pushed to the stack.
// Function results are returned via stack as well.
//
// For the sake of efficiency, it stores different types separately.
// If int was pushed with PushInt(), it should be retrieved by PopInt().
// It's a bad idea to do a Push() and then PopInt() and vice-versa.
type ValueStack struct {
	objects     []interface{}
	ints        []int
	variadicLen int
}

// Reset empties the stack.
func (s *ValueStack) Reset() { _ = "STUB: not implemented"; return }

// Pop removes the top stack element and returns it.
// Important: for int-typed values, use PopInt.
func (s *ValueStack) Pop() interface{} { _ = "STUB: not implemented"; return nil }

// PopInt removes the top stack element and returns it.
func (s *ValueStack) PopInt() int { _ = "STUB: not implemented"; return 0 }

// PopVariadic removes the `...` argument and returns it as a slice.
//
// Slice elements are in the order they were passed to the function,
// for example, a call Sprintf("%s:%d", filename, line) returns
// the slice []interface{filename, line}.
func (s *ValueStack) PopVariadic() []interface{} { _ = "STUB: not implemented"; return nil }

// Push adds x to the stack.
// Important: for int-typed values, use PushInt.
func (s *ValueStack) Push(x interface{}) { _ = "STUB: not implemented"; return }

// PushInt adds x to the stack.
func (s *ValueStack) PushInt(x int) { _ = "STUB: not implemented"; return }
