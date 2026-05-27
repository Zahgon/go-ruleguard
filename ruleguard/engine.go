package ruleguard

import (
	"go/ast"
	"go/build"
	"go/types"
	"io"
	"sync"

	"github.com/quasilyte/go-ruleguard/ruleguard/ir"
	"github.com/quasilyte/go-ruleguard/ruleguard/quasigo"
)

type engine struct {
	state *engineState

	ruleSet *goRuleSet
}

func newEngine() *engine { _ = "STUB: not implemented"; return nil }

func (e *engine) LoadedGroups() []GoRuleGroup { _ = "STUB: not implemented"; return nil }

func (e *engine) Load(ctx *LoadContext, buildContext *build.Context, filename string, r io.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *engine) LoadFromIR(ctx *LoadContext, buildContext *build.Context, filename string, f *ir.File) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *engine) Run(ctx *RunContext, buildContext *build.Context, f *ast.File) error {
	_ = "STUB: not implemented"
	return nil
}

// engineState is a shared state inside the engine.
// Its access is synchronized, unlike the RunnerState which should be thread-local.
type engineState struct {
	env *quasigo.Env

	typeByFQNMu sync.RWMutex
	typeByFQN   map[string]types.Type

	pkgCacheMu sync.RWMutex
	// pkgCache contains all imported packages, from any importer.
	pkgCache map[string]*types.Package
}

func newEngineState() *engineState { _ = "STUB: not implemented"; return nil }

func (state *engineState) GetCachedPackage(pkgPath string) *types.Package {
	_ = "STUB: not implemented"
	return nil
}

func (state *engineState) AddCachedPackage(pkgPath string, pkg *types.Package) {
	_ = "STUB: not implemented"
	return
}

func (state *engineState) addCachedPackage(pkgPath string, pkg *types.Package) {
	_ = "STUB: not implemented"
	return
}

// Also add all complete packages that are dependencies of the pkg.
// This way we cache more and avoid duplicated package loading
// which can lead to typechecking issues.
//
// Note that it does not increase our memory consumption
// as these packages are reachable via pkg, so they'll
// not be freed by GC anyway.

func (state *engineState) FindType(importer *goImporter, currentPkg *types.Package, fqn string) (types.Type, error) {
	_ = "STUB: not implemented"
	// TODO(quasilyte): we can pre-populate the cache during the Load() phase.
	// If we inspect the AST of a user function, all constant FQN can be preloaded.
	// It could be a good thing as Load() is not expected to be executed in
	// concurrent environment, so write-locking is not a big deal there.
	return *new(types.Type), nil
}

// Code below is under a write critical section.

func (state *engineState) findTypeNoCache(importer *goImporter, currentPkg *types.Package, fqn string) (types.Type, error) {
	_ = "STUB: not implemented"
	return *new(types.Type), nil
}

func inferBuildContext() *build.Context {
	_ = "STUB: not implemented"
	// Inherit most fields from the build.Default.
	return nil
}
