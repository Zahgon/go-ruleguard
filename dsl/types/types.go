// Package types mimics the https://golang.org/pkg/go/types/ package.
// It also contains some extra utility functions, they're defined in ext.go file.
package types

// Implements reports whether a given type implements the specified interface.
func Implements(typ Type, iface *Interface) bool {
	_ = "STUB: not implemented"

	// Identical reports whether x and y are identical types. Receivers of Signature types are ignored.
	return false
}

func Identical(x, y Type) bool {
	_ = "STUB: not implemented"

	// A Type represents a type of Go. All types implement the Type interface.
	return false
}

type Type interface {
	// Underlying returns the underlying type of a type.
	Underlying() Type

	// String returns a string representation of a type.
	String() string
}

type (
	// An Array represents an array type.
	Array struct{}

	// A Slice represents a slice type.
	Slice struct{}

	// A Pointer represents a pointer type.
	Pointer struct{}

	// An Interface represents an interface type.
	Interface struct{}

	// A struct represents a struct type.
	Struct struct{}
)

// NewArray returns a new array type for the given element type and length.
// A negative length indicates an unknown length.
func NewArray(elem Type, len int) *Array {
	_ = "STUB: not implemented"

	// Elem returns element type of array.
	return nil
}

func (*Array) Elem() Type {
	_ = "STUB: not implemented"

	// NewSlice returns a new slice type for the given element type.
	return *new(Type)
}

func NewSlice(elem Type) *Slice {
	_ = "STUB: not implemented"

	// Elem returns element type of slice.
	return nil
}

func (*Slice) Elem() Type {
	_ = "STUB: not implemented"

	// Len returns the length of array.
	// A negative result indicates an unknown length.
	return *new(Type)
}

func (*Array) Len() int {
	_ = "STUB: not implemented"

	// NewPointer returns a new pointer type for the given element (base) type.
	return 0
}

func NewPointer(elem Type) *Pointer {
	_ = "STUB: not implemented"

	// Elem returns the element type for the given pointer.
	return nil
}

func (*Pointer) Elem() Type { _ = "STUB: not implemented"; return *new(Type) }

func (*Struct) NumFields() int { _ = "STUB: not implemented"; return 0 }

func (*Struct) Field(i int) *Var { _ = "STUB: not implemented"; return nil }

type Var struct{}

func (*Var) Embedded() bool { _ = "STUB: not implemented"; return false }

func (*Var) Type() Type { _ = "STUB: not implemented"; return *new(Type) }
