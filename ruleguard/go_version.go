package ruleguard

import (
	"go/token"
)

type GoVersion struct {
	Major int
	Minor int
}

func (ver GoVersion) IsAny() bool { _ = "STUB: not implemented"; return false }

func ParseGoVersion(version string) (GoVersion, error) {
	_ = "STUB: not implemented"
	return *new(GoVersion), nil
}

func versionCompare(x GoVersion, op token.Token, y GoVersion) bool {
	_ = "STUB: not implemented"
	return false

	// ==
}

// !=

// >

// >=

// <

// <=
