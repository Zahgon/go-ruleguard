package xsrcimporter

import (
	"go/build"
	"go/token"
	"go/types"
	"unsafe"
)

func New(ctxt *build.Context, fset *token.FileSet) types.Importer {
	_ = "STUB: not implemented"
	return *new(types.Importer)
}

type iface struct {
	_    *byte
	data unsafe.Pointer
}

type srcImporter struct {
	ctxt *build.Context
	_    *token.FileSet
	_    types.Sizes
	_    map[string]*types.Package
}
