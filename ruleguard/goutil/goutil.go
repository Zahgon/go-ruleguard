package goutil

import (
	"go/ast"
	"go/token"
	"go/types"
)

// SprintNode returns the textual representation of n.
// If fset is nil, freshly created file set will be used.
func SprintNode(fset *token.FileSet, n ast.Node) string { _ = "STUB: not implemented"; return "" }

type LoadConfig struct {
	Fset     *token.FileSet
	Filename string
	Data     interface{}
	Importer types.Importer
}

type LoadResult struct {
	Pkg    *types.Package
	Types  *types.Info
	Syntax *ast.File
}

func LoadGoFile(config LoadConfig) (*LoadResult, error) { _ = "STUB: not implemented"; return nil, nil }
