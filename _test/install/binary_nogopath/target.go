package target

func add(x, y int) int { _ = "STUB: not implemented"; return 0 }

func test(b bool) { _ = "STUB: not implemented"; return }

// Not fmt.Stringer
// Not addressable

type Foo struct{}

func (*Foo) String() string { _ = "STUB: not implemented"; return "" }
