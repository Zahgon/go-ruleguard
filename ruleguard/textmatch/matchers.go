package textmatch

// inputValue is a wrapper for string|[]byte.
//
// We hold both values to avoid string->[]byte and vice versa
// conversions when doing Match and MatchString.
type inputValue struct {
	s string
	b []byte
}

func newInputValue(s string) inputValue { _ = "STUB: not implemented"; return *new(inputValue) }

type containsLiteralMatcher struct{ value inputValue }

func (m *containsLiteralMatcher) MatchString(s string) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *containsLiteralMatcher) Match(b []byte) bool { _ = "STUB: not implemented"; return false }

type prefixLiteralMatcher struct{ value inputValue }

func (m *prefixLiteralMatcher) MatchString(s string) bool { _ = "STUB: not implemented"; return false }

func (m *prefixLiteralMatcher) Match(b []byte) bool { _ = "STUB: not implemented"; return false }

type suffixLiteralMatcher struct{ value inputValue }

func (m *suffixLiteralMatcher) MatchString(s string) bool { _ = "STUB: not implemented"; return false }

func (m *suffixLiteralMatcher) Match(b []byte) bool { _ = "STUB: not implemented"; return false }

type eqLiteralMatcher struct{ value inputValue }

func (m *eqLiteralMatcher) MatchString(s string) bool { _ = "STUB: not implemented"; return false }

func (m *eqLiteralMatcher) Match(b []byte) bool { _ = "STUB: not implemented"; return false }

type prefixRunePredMatcher struct{ pred func(rune) bool }

func (m *prefixRunePredMatcher) MatchString(s string) bool { _ = "STUB: not implemented"; return false }

func (m *prefixRunePredMatcher) Match(b []byte) bool { _ = "STUB: not implemented"; return false }
