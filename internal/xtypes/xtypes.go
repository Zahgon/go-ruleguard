package xtypes

import (
	"go/types"
)

// Implements reports whether type v implements iface.
//
// Unlike types.Implements(), it permits X and Y named types
// to be considered identical even if their addresses are different.
func Implements(v types.Type, iface *types.Interface) bool { _ = "STUB: not implemented"; return false }

// A concrete type v implements iface if it implements all methods of iface.

// Identical reports whether x and y are identical types.
//
// Unlike types.Identical(), it permits X and Y named types
// to be considered identical even if their addresses are different.
func Identical(x, y types.Type) bool { _ = "STUB: not implemented"; return false }

func typeIdentical(x, y types.Type, p *ifacePair) bool { _ = "STUB: not implemented"; return false }

// Basic types are singletons except for the rune and byte
// aliases, thus we cannot solely rely on the x == y check
// above. See also comment in TypeName.IsAlias.

// Two array types are identical if they have identical element types
// and the same array length.

// If one or both array lengths are unknown (< 0) due to some error,
// assume they are the same to avoid spurious follow-on errors.

// Two slice types are identical if they have identical element types.

// Two struct types are identical if they have the same sequence of fields,
// and if corresponding fields have the same names, and identical types,
// and identical tags. Two embedded fields are considered to have the same
// name. Lower-case field names from different packages are always different.

// Two pointer types are identical if they have identical base types.

// Two tuples types are identical if they have the same number of elements
// and corresponding elements have identical types.

// Two function types are identical if they have the same number of parameters
// and result values, corresponding parameter and result types are identical,
// and either both functions are variadic or neither is. Parameter and result
// names are not required to match.

// TODO(quasilyte): do we want to match generic union types too?
// It would require copying a lot of code from the go/types.

// Two interface types are identical if they have the same set of methods with
// the same names and identical function types. Lower-case method names from
// different packages are always different. The order of the methods is irrelevant.

// Interface types are the only types where cycles can occur
// that are not "terminated" via named types; and such cycles
// can only be created via method parameter types that are
// anonymous interfaces (directly or indirectly) embedding
// the current interface. Example:
//
//    type T interface {
//        m() interface{T}
//    }
//
// If two such (differently named) interfaces are compared,
// endless recursion occurs if the cycle is not detected.
//
// If x and y were compared before, they must be equal
// (if they were not, the recursion would have stopped);
// search the ifacePair stack for the same pair.
//
// This is a quadratic algorithm, but in practice these stacks
// are extremely short (bounded by the nesting depth of interface
// type declarations that recur via parameter types, an extremely
// rare occurrence). An alternative implementation might use a
// "visited" map, but that is probably less efficient overall.

// same pair was compared before

// Two map types are identical if they have identical key and value types.

// Two channel types are identical if they have identical value types
// and the same direction.

// Two named types are identical if their type names originate
// in the same type declaration.

// nothing to do (x and y being equal is caught in the very beginning of this function)

// an alias type is identical if the type it's an alias of is identical to it.

// avoid a crash in case of nil type

// An ifacePair is a node in a stack of interface type pairs compared for identity.
type ifacePair struct {
	x    *types.Interface
	y    *types.Interface
	prev *ifacePair
}

func (p *ifacePair) identical(q *ifacePair) bool { _ = "STUB: not implemented"; return false }

func sameID(obj types.Object, pkg *types.Package, name string) bool {
	_ = "STUB: not implemented"
	// spec:
	// "Two identifiers are different if they are spelled differently,
	// or if they appear in different packages and are not exported.
	// Otherwise, they are the same."
	return false
}

// obj.Name == name

// not exported, so packages must be the same (pkg == nil for
// fields in Universe scope; this can only happen for types
// introduced via Eval)

// pkg != nil && obj.pkg != nil
