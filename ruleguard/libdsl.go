package ruleguard

import (
	"github.com/quasilyte/go-ruleguard/ruleguard/quasigo"
)

// This file implements `dsl/*` packages as native functions in quasigo.
//
// Every function and method defined in any `dsl/*` package should have
// associated Go function that implements it.
//
// In quasigo, it's impossible to have a pointer to an interface and
// non-pointer struct type. All interface type methods have FQN without `*` prefix
// while all struct type methods always begin with `*`.
//
// Fields are readonly.
// Field access is compiled into a method call that have a name identical to the field.
// For example, `foo.Bar` field access will be compiled as `foo.Bar()`.
// This may change in the future; benchmarks are needed to figure out
// what is more efficient: reflect-based field access or a function call.
//
// To keep this code organized, every type and package functions are represented
// as structs with methods. Then we bind a method value to quasigo symbol.
// The naming scheme is `dsl{$name}Package` for packages and `dsl{$pkg}{$name}` for types.

func initEnv(state *engineState, env *quasigo.Env) { _ = "STUB: not implemented"; return }

type quasigoNative interface {
	funcs() map[string]func(*quasigo.ValueStack)
}

type dslTypesType struct{}

func (native dslTypesType) funcs() map[string]func(*quasigo.ValueStack) {
	_ = "STUB: not implemented"
	return nil
}

func (dslTypesType) Underlying(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

func (dslTypesType) String(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

type dslTypesInterface struct{}

func (native dslTypesInterface) funcs() map[string]func(*quasigo.ValueStack) {
	_ = "STUB: not implemented"
	return nil
}

func (dslTypesInterface) Underlying(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

func (dslTypesInterface) String(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

type dslTypesSlice struct{}

func (native dslTypesSlice) funcs() map[string]func(*quasigo.ValueStack) {
	_ = "STUB: not implemented"
	return nil
}

func (dslTypesSlice) Underlying(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

func (dslTypesSlice) String(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

func (dslTypesSlice) Elem(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

type dslTypesArray struct{}

func (native dslTypesArray) funcs() map[string]func(*quasigo.ValueStack) {
	_ = "STUB: not implemented"
	return nil
}

func (dslTypesArray) Underlying(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

func (dslTypesArray) String(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

func (dslTypesArray) Elem(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

func (dslTypesArray) Len(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

type dslTypesPointer struct{}

func (native dslTypesPointer) funcs() map[string]func(*quasigo.ValueStack) {
	_ = "STUB: not implemented"
	return nil
}

func (dslTypesPointer) Underlying(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

func (dslTypesPointer) String(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

func (dslTypesPointer) Elem(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

type dslTypesStruct struct{}

func (native dslTypesStruct) funcs() map[string]func(*quasigo.ValueStack) {
	_ = "STUB: not implemented"
	return nil
}

func (dslTypesStruct) Underlying(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

func (dslTypesStruct) String(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

func (dslTypesStruct) NumFields(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

func (dslTypesStruct) Field(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

type dslTypesPackage struct{}

func (native dslTypesPackage) funcs() map[string]func(*quasigo.ValueStack) {
	_ = "STUB: not implemented"
	return nil
}

func (dslTypesPackage) Implements(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

func (dslTypesPackage) Identical(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

func (dslTypesPackage) NewArray(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

func (dslTypesPackage) NewSlice(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

func (dslTypesPackage) NewPointer(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

func (dslTypesPackage) AsArray(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

func (dslTypesPackage) AsSlice(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

func (dslTypesPackage) AsPointer(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

func (dslTypesPackage) AsInterface(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

func (dslTypesPackage) AsStruct(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

type dslTypesVar struct{}

func (native dslTypesVar) funcs() map[string]func(*quasigo.ValueStack) {
	_ = "STUB: not implemented"
	return nil
}

func (dslTypesVar) Embedded(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

func (dslTypesVar) Type(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

type dslDoContext struct{}

func (native dslDoContext) funcs() map[string]func(*quasigo.ValueStack) {
	_ = "STUB: not implemented"
	return nil
}

func (native dslDoContext) Var(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

func (native dslDoContext) SetReport(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

func (native dslDoContext) SetSuggest(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

type dslMatchedText struct{}

func (native dslMatchedText) funcs() map[string]func(*quasigo.ValueStack) {
	_ = "STUB: not implemented"
	return nil
}

func (dslMatchedText) String(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

type dslDoVarRepr struct {
	params *filterParams
	name   string
}

type dslDoVar struct{}

func (native dslDoVar) funcs() map[string]func(*quasigo.ValueStack) {
	_ = "STUB: not implemented"
	return nil
}

func (dslDoVar) Text(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

func (dslDoVar) Type(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

type dslVarFilterContext struct {
	state *engineState
}

func (native dslVarFilterContext) funcs() map[string]func(*quasigo.ValueStack) {
	_ = "STUB: not implemented"
	return nil
}

func (dslVarFilterContext) Type(stack *quasigo.ValueStack) { _ = "STUB: not implemented"; return }

func (native dslVarFilterContext) SizeOf(stack *quasigo.ValueStack) {
	_ = "STUB: not implemented"
	return
}

func (native dslVarFilterContext) GetType(stack *quasigo.ValueStack) {
	_ = "STUB: not implemented"
	return
}

func (native dslVarFilterContext) GetInterface(stack *quasigo.ValueStack) {
	_ = "STUB: not implemented"
	return
}

// Not found or not an interface
