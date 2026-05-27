package types

// Method stubs to make various types implement Type interface.
//
// Nothing interesting here, hence it's moved to a separate file.

func (*Array) String() string     { _ = "STUB: not implemented"; return "" }
func (*Slice) String() string     { _ = "STUB: not implemented"; return "" }
func (*Pointer) String() string   { _ = "STUB: not implemented"; return "" }
func (*Interface) String() string { _ = "STUB: not implemented"; return "" }
func (*Struct) String() string    { _ = "STUB: not implemented"; return "" }

func (*Array) Underlying() Type     { _ = "STUB: not implemented"; return *new(Type) }
func (*Slice) Underlying() Type     { _ = "STUB: not implemented"; return *new(Type) }
func (*Pointer) Underlying() Type   { _ = "STUB: not implemented"; return *new(Type) }
func (*Interface) Underlying() Type { _ = "STUB: not implemented"; return *new(Type) }
func (*Struct) Underlying() Type    { _ = "STUB: not implemented"; return *new(Type) }
