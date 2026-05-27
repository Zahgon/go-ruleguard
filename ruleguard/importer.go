package ruleguard

import (
	"go/build"
	"go/token"
	"go/types"
)

// goImporter is a `types.Importer` that tries to load a package no matter what.
// It iterates through multiple import strategies and accepts whatever succeeds first.
type goImporter struct {
	// TODO(quasilyte): share importers with gogrep?

	state *engineState

	defaultImporter types.Importer
	srcImporter     types.Importer

	fset         *token.FileSet
	buildContext *build.Context

	debugImports bool
	debugPrint   func(string)
}

type goImporterConfig struct {
	fset         *token.FileSet
	debugImports bool
	debugPrint   func(string)
	buildContext *build.Context
}

func newGoImporter(state *engineState, config goImporterConfig) *goImporter {
	_ = "STUB: not implemented"
	return nil
}

func (imp *goImporter) Import(path string) (*types.Package, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (imp *goImporter) initSourceImporter() { _ = "STUB: not implemented"; return }
