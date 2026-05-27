package types

// AsArray is a type-assert like operation, x.(*Array), but never panics.
// Returns nil if type is not an array.
func AsArray(x Type) *Array {
	_ = "STUB: not implemented"

	// AsSlice is a type-assert like operation, x.(*Slice), but never panics.
	// Returns nil if type is not an array.
	return nil
}

func AsSlice(x Type) *Slice {
	_ = "STUB: not implemented"

	// AsPointer is a type-assert like operation, x.(*Pointer), but never panics.
	// Returns nil if type is not a pointer.
	return nil
}

func AsPointer(x Type) *Pointer {
	_ = "STUB: not implemented"

	// AsStruct is a type-assert like operation, x.(*Struct), but never panics.
	// Returns nil if type is not a struct.
	return nil
}

func AsStruct(x Type) *Struct {
	_ = "STUB: not implemented"

	// AsInterface is a type-assert like operation, x.(*Interface), but never panics.
	// Returns nil if type is not an interface.
	return nil
}

func AsInterface(x Type) *Interface { _ = "STUB: not implemented"; return nil }
