package dsl

// Matcher is a main API group-level entry point.
// It's used to define and configure the group rules.
// It also represents a map of all rule-local variables.
type Matcher map[string]Var

// Import loads given package path into a rule group imports table.
//
// That table is used during the rules compilation.
//
// The table has the following effect on the rules:
//   - For type expressions, it's used to resolve the
//     full package paths of qualified types, like `foo.Bar`.
//     If Import(`a/b/foo`) is called, `foo.Bar` will match
//     `a/b/foo.Bar` type during the pattern execution.
func (m Matcher) Import(pkgPath string) {
	_ = "STUB: not implemented"

	// ImportAs is like Import, but can handle "/v2" packages
	// and package name conflicts (e.g. "x/path" vs "y/path").
	return
}

func (m Matcher) ImportAs(pkgPath, localName string) {
	_ = "STUB: not implemented"

	// Match specifies a set of patterns that match a rule being defined.
	// Pattern matching succeeds if at least 1 pattern matches.
	//
	// If none of the given patterns matched, rule execution stops.
	return
}

func (m Matcher) Match(pattern string, alternatives ...string) Matcher {
	_ = "STUB: not implemented"

	// MatchComment is like Match, but handles only comments and uses regexp patterns.
	//
	// Multi-line /**/ comments are passed as a single string.
	// Single-line // comments are passed line-by-line.
	//
	// Hint: if you want to match a plain text and don't want to do meta char escaping,
	// prepend `\Q` to your pattern. `\Qf(x)` will match `f(x)` as a plain text
	// and there is no need to escape the `(` and `)` chars.
	//
	// Named regexp capture groups can be accessed using the usual indexing notation.
	//
	// Given this pattern:
	//
	//	`(?P<first>\d+)\.(\d+).(?P<second>\d+)`
	//
	// And this input comment: `// 14.6.600`
	//
	// We'll get these submatches:
	//
	//	m["$$"] => `14.6.600`
	//	m["first"] => `14`
	//	m["second"] => `600`
	//
	// All usual filters can be applied:
	//
	//	Where(!m["first"].Text.Matches(`foo`))
	//
	// You can use this to reject some matches (allow-list behavior).
	return *new(Matcher)
}

func (m Matcher) MatchComment(pattern string, alternatives ...string) Matcher {
	_ = "STUB: not implemented"

	// Where applies additional constraint to a match.
	// If a given cond is not satisfied, a match is rejected and
	// rule execution stops.
	return *new(Matcher)
}

func (m Matcher) Where(cond bool) Matcher {
	_ = "STUB: not implemented"

	// Report prints a message if associated rule match is successful.
	//
	// A message is a string that can contain interpolated expressions.
	// For every matched variable it's possible to interpolate
	// their printed representation into the message text with $<name>.
	// An entire match can be addressed with $$.
	return *new(Matcher)
}

func (m Matcher) Report(message string) Matcher {
	_ = "STUB: not implemented"

	// Suggest assigns a quickfix suggestion for the matched code.
	return *new(Matcher)
}

func (m Matcher) Suggest(suggestion string) Matcher {
	_ = "STUB: not implemented"
	return *new(Matcher)
}

func (m Matcher) Do(fn func(*DoContext)) Matcher {
	_ = "STUB: not implemented"

	// At binds the reported node to a named submatch.
	// If no explicit location is given, the outermost node ($$) is used.
	return *new(Matcher)
}

func (m Matcher) At(v Var) Matcher {
	_ = "STUB: not implemented"

	// File returns the current file context.
	return *new(Matcher)
}

func (m Matcher) File() File {
	_ = "STUB: not implemented"

	// GoVersion returns the analyzer associated target Go language version.
	return *new(File)
}

func (m Matcher) GoVersion() GoVersion {
	_ = "STUB: not implemented"

	// Deadcode reports whether this match is contained inside a dead code path.
	return *new(GoVersion)
}

func (m Matcher) Deadcode() bool {
	_ = "STUB: not implemented"

	// Var is a pattern variable that describes a named submatch.
	return false
}

type Var struct {
	// Pure reports whether expr matched by var is side-effect-free.
	Pure bool

	// Const reports whether expr matched by var is a constant value.
	Const bool

	// ConstSlice reports whether expr matched by var is a slice literal
	// consisting of constant elements.
	//
	// We need a separate Const-like predicate here because Go doesn't
	// treat slices of const elements as constants, so including
	// them in Const would be incorrect.
	// Use `m["x"].Const || m["x"].ConstSlice` when you need
	// to have extended definition of "constant value".
	//
	// Some examples:
	//     []byte("foo") -- constant byte slice
	//     []byte{'f', 'o', 'o'} -- same constant byte slice
	//     []int{1, 2} -- constant int slice
	ConstSlice bool

	// Value is a compile-time computable value of the expression.
	Value ExprValue

	// Addressable reports whether the corresponding expression is addressable.
	// See https://golang.org/ref/spec#Address_operators.
	Addressable bool

	// Comparable reports whether the corresponding expression value is comparable.
	// See https://pkg.go.dev/go/types#Comparable.
	Comparable bool

	// Type is a type of a matched expr.
	//
	// For function call expressions, a type is a function result type,
	// but for a function expression itself it's a *types.Signature.
	//
	// Suppose we have a `a.b()` expression:
	//	`$x()` m["x"].Type is `a.b` function type
	//	`$x` m["x"].Type is `a.b()` function call result type
	Type ExprType

	SinkType SinkType

	// Object is an associated "go/types" Object.
	Object TypesObject

	// Text is a captured node text as in the source code.
	Text MatchedText

	// Node is a captured AST node.
	Node MatchedNode

	// Line is a source code line number that contains this match.
	// If this match is multi-line, this is the first line number.
	Line int
}

// Filter applies a custom predicate function on a submatch.
//
// The callback function should use VarFilterContext to access the
// information that is usually accessed through Var.
// For example, `VarFilterContext.Type` is mapped to `Var.Type`.
func (Var) Filter(pred func(*VarFilterContext) bool) bool {
	_ = "STUB: not implemented"

	// Contains runs a sub-search from a given pattern using the captured
	// vars from the original pattern match.
	//
	// For example, given the Match(`$lhs = append($lhs, $x)`) pattern,
	// we can do m["lhs"].Contains(`$x`) and learn whether $lhs contains
	// $x as its sub-expression.
	//
	// Experimental: this function is not part of the stable API.
	return false
}

func (Var) Contains(pattern string) bool {
	_ = "STUB: not implemented"

	// MatchedNode represents an AST node associated with a named submatch.
	return false
}

type MatchedNode struct{}

// Is reports whether a matched node AST type is compatible with the specified type.
// A valid argument is a ast.Node implementing type name from the "go/ast" package.
// Examples: "BasicLit", "Expr", "Stmt", "Ident", "ParenExpr".
// See https://golang.org/pkg/go/ast/.
func (MatchedNode) Is(typ string) bool {
	_ = "STUB: not implemented"

	// Parent returns a matched node parent.
	return false
}

func (MatchedNode) Parent() Node {
	_ = "STUB: not implemented"

	// Node represents an AST node somewhere inside a match.
	// Unlike MatchedNode, it doesn't have to be associated with a named submatch.
	return *new(Node)
}

type Node struct{}

// Is reports whether a node AST type is compatible with the specified type.
// See `MatchedNode.Is` for the full reference.
func (Node) Is(typ string) bool {
	_ = "STUB: not implemented"

	// ExprValue describes a compile-time computable value of a matched expr.
	return false
}

type ExprValue struct{}

// Int returns compile-time computable int value of the expression.
// If value can't be computed, condition will fail.
func (ExprValue) Int() int {
	_ = "STUB: not implemented"

	// TypesObject is a types.Object mapping.
	return 0
}

type TypesObject struct{}

// Is reports whether an associated types.Object is compatible with the specified type.
// A valid argument is a types.Object type name from the "go/types" package.
// Examples: "Func", "Var", "Const", "TypeName", "Label", "PkgName", "Builtin", "Nil"
// See https://golang.org/pkg/go/types/.
func (TypesObject) Is(typ string) bool {
	_ = "STUB: not implemented"

	// IsGlobal reports whether an associated types.Object is defined in global scope.
	return false
}

func (TypesObject) IsGlobal() bool {
	_ = "STUB: not implemented"

	// IsVariadicParam reports whether this object represents a function variadic param.
	// This property is not propagated between the assignments.
	return false
}

func (TypesObject) IsVariadicParam() bool { _ = "STUB: not implemented"; return false }

type SinkType struct{}

// Is reports whether a type is identical to a given type.
// Works like ExprType.Is method.
func (SinkType) Is(typ string) bool {
	_ = "STUB: not implemented"

	// ExprType describes a type of a matcher expr.
	return false
}

type ExprType struct {
	// Size represents expression type size in bytes.
	//
	// For expressions of unknown size, like type params in generics,
	// any filter using this operand will fail.
	Size int
}

// IdenticalTo applies types.Identical(this, v.Type) operation.
// See https://golang.org/pkg/go/types/#Identical function documentation.
//
// Experimental: this function is not part of the stable API.
func (ExprType) IdenticalTo(v Var) bool {
	_ = "STUB: not implemented"

	// Underlying returns expression type underlying type.
	// See https://golang.org/pkg/go/types/#Type Underlying() method documentation.
	// Read https://golang.org/ref/spec#Types section to learn more about underlying types.
	return false
}

func (ExprType) Underlying() ExprType {
	_ = "STUB: not implemented"
	return *

	// AssignableTo reports whether a type is assign-compatible with a given type.
	// See https://golang.org/pkg/go/types/#AssignableTo.
	new(ExprType)
}

func (ExprType) AssignableTo(typ string) bool {
	_ = "STUB: not implemented"

	// ConvertibleTo reports whether a type is conversible to a given type.
	// See https://golang.org/pkg/go/types/#ConvertibleTo.
	return false
}

func (ExprType) ConvertibleTo(typ string) bool {
	_ = "STUB: not implemented"

	// Implements reports whether a type implements a given interface.
	// See https://golang.org/pkg/go/types/#Implements.
	return false
}

func (ExprType) Implements(typ typeName) bool {
	_ = "STUB: not implemented"

	// HasMethod reports whether a type has a given method.
	// Unlike Implements(), it will work for both value and pointer types.
	//
	// fn argument is a function signature, like `WriteString(string) (int, error)`.
	// It can also be in form of a method reference for importable types: `io.StringWriter.WriteString`.
	//
	// To avoid confusion with Implements() method, here is a hint when to use which:
	//
	//   - To check if it's possible to call F on x, use HasMethod(F)
	//   - To check if x can be passed as I interface, use Implements(I)
	return false
}

func (ExprType) HasMethod(fn string) bool {
	_ = "STUB: not implemented"

	// Is reports whether a type is identical to a given type.
	return false
}

func (ExprType) Is(typ string) bool {
	_ = "STUB: not implemented"

	// HasPointers reports whether a type contains at least one pointer.
	//
	// We try to be as close to the Go sense of pointer-free objects as possible,
	// therefore string type is not considered to be a pointer-free type.
	//
	// This function may return "true" for some complicated cases as a
	// conservative result. It never returns "false" for a type that
	// actually contains a pointer.
	//
	// So this function is mostly useful for !HasPointers() form.
	return false
}

func (ExprType) HasPointers() bool {
	_ = "STUB: not implemented"

	// OfKind reports whether a matched expr type is compatible with the specified kind.
	//
	// Only a few "kinds" are recognized, the list is provided below.
	//
	//		"integer"  -- typ is *types.Basic, where typ.Info()&types.Integer != 0
	//		"unsigned" -- typ is *types.Basic, where typ.Info()&types.Unsigned != 0
	//		"float"    -- typ is *types.Basic, where typ.Info()&types.Float != 0
	//		"complex"  -- typ is *types.Basic, where typ.Info()&types.Complex != 0
	//		"untyped"  -- typ is *types.Basic, where typ.Info()&types.Untyped != 0
	//		"numeric"  -- typ is *types.Basic, where typ.Info()&types.Numeric != 0
	//	 "signed"   -- identical to `OfKind("integer") && !OfKind("unsigned")`
	//	 "int"      -- int, int8, int16, int32, int64
	//	 "uint"     -- uint, uint8, uint16, uint32, uint64
	//
	// Note: "int" will include "rune" as well, as it's an alias.
	// In the same manner, "uint" includes the "byte" type.
	//
	// Using OfKind("unsigned") is more efficient (and concise) than using a set
	// of or-conditions with Is("uint8"), Is("uint16") and so on.
	return false
}

func (ExprType) OfKind(kind string) bool {
	_ = "STUB: not implemented"

	// MatchedText represents a source text associated with a matched node.
	return false
}

type MatchedText string

// Matches reports whether the text matches the given regexp pattern.
func (MatchedText) Matches(pattern string) bool {
	_ = "STUB: not implemented"

	// String represents an arbitrary string-typed data.
	return false
}

type String string

// Matches reports whether a string matches the given regexp pattern.
func (String) Matches(pattern string) bool {
	_ = "STUB: not implemented"

	// File represents the current Go source file.
	return false
}

type File struct {
	// Name is a file base name.
	Name String

	// PkgPath is a file package path.
	// Examples: "io/ioutil", "strings", "github.com/quasilyte/go-ruleguard/dsl".
	PkgPath String
}

// Imports reports whether the current file imports the given path.
func (File) Imports(path string) bool {
	_ = "STUB: not implemented"

	// GoVersion is an analysis target go language version.
	// It can be compared to Go versions like "1.10", "1.16" using
	// the associated methods.
	return false
}

type GoVersion struct{}

// Eq asserts that target Go version is equal to (==) specified version.
func (GoVersion) Eq(version string) bool {
	_ = "STUB: not implemented"

	// GreaterEqThan asserts that target Go version is greater or equal than (>=) specified version.
	return false
}

func (GoVersion) GreaterEqThan(version string) bool {
	_ = "STUB: not implemented"

	// GreaterThan asserts that target Go version is greater than (>) specified version.
	return false
}

func (GoVersion) GreaterThan(version string) bool {
	_ = "STUB: not implemented"

	// LessThan asserts that target Go version is less than (<) specified version.
	return false
}

func (GoVersion) LessThan(version string) bool {
	_ = "STUB: not implemented"

	// LessEqThan asserts that target Go version is less or equal than (<=) specified version.
	return false
}

func (GoVersion) LessEqThan(version string) bool {
	_ = "STUB: not implemented"

	// typeName is a helper type used to document function params better.
	//
	// A type name can be:
	//   - builtin type name: `error`, `string`, etc.
	//   - qualified name from a standard library: `io.Reader`, etc.
	//   - fully-qualified type name, like `github.com/username/pkgname.TypeName`
	//
	// typeName is also affected by a local import table, which can override
	// how qualified names are interpreted.
	// See `Matcher.Import` for more info.
	return false
}

type typeName = string
