package dsl

import (
	"github.com/quasilyte/go-ruleguard/dsl/types"
)

// VarFilterContext carries Var and environment information into the filter function.
// It's an input parameter type for the Var.Filter function callback.
type VarFilterContext struct {
	// Type is mapped to Var.Type field.
	Type types.Type
}

// SizeOf returns the size of the given type.
// It uses the ruleguard.Context.Sizes to calculate the result.
func (*VarFilterContext) SizeOf(x types.Type) int {
	_ = "STUB: not implemented"

	// GetType finds a type value by a given name.
	// If a type can't be found (or a name is malformed), this function panics.
	return 0
}

func (*VarFilterContext) GetType(name typeName) types.Type {
	_ = "STUB: not implemented"

	// GetInterface finds a type value that represents an interface by a given name.
	// Works like `types.AsInterface(ctx.GetType(name))`.
	return *new(types.Type)
}

func (*VarFilterContext) GetInterface(name typeName) *types.Interface {
	_ = "STUB: not implemented"
	return nil
}
