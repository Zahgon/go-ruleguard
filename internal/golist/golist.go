package golist

// Package is `go list --json` output structure.
type Package struct {
	Dir        string   // directory containing package sources
	ImportPath string   // import path of package in dir
	GoFiles    []string // .go source files (excluding CgoFiles, TestGoFiles, XTestGoFiles)
}

// JSON runs `go list --json` for the specified pkgName and returns the parsed JSON.
func JSON(pkgPath string) (*Package, error) { _ = "STUB: not implemented"; return nil, nil }
