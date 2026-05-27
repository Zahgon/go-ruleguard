package textmatch

import (
	"regexp/syntax"
)

func compile(s string) (Pattern, error) { _ = "STUB: not implemented"; return *new(Pattern), nil }

func compileOptimized(s string, re *syntax.Regexp) Pattern {
	_ = "STUB: not implemented"
	// .*
	return *new(Pattern)
}

// "literal"

// ^

// $

// TODO: analyze what kind of regexps people use in rules
// more often and optimize those as well.

// lit => strings.Contains($input, lit)

// `.*` lit `.*` => strings.Contains($input, lit)

// `^` lit => strings.HasPrefix($input, lit)

// lit `$` => strings.HasSuffix($input, lit)

// `^` lit `$` => $input == lit

// `^\p{Lu}` => prefixRunePredMatcher:unicode.IsUpper
// `^\p{Ll}` => prefixRunePredMatcher:unicode.IsLower

// Can't optimize.
