package textmatch

// Pattern is a compiled regular expression.
type Pattern interface {
	MatchString(s string) bool
	Match(b []byte) bool
}

// Compile parses a regular expression and returns a compiled
// pattern that can match inputs described by the regexp.
//
// Semantically it's close to the regexp.Compile, but
// it does recognize some common patterns and creates
// a more optimized matcher for them.
func Compile(re string) (Pattern, error) {
	_ = "STUB: not implemented"

	// IsRegexp reports whether p is implemented using regexp.
	// False means that the underlying matcher is something optimized.
	return *new(Pattern), nil
}

func IsRegexp(p Pattern) bool { _ = "STUB: not implemented"; return false }
