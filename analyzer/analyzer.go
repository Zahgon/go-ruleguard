package analyzer

import (
	"sync"

	"golang.org/x/tools/go/analysis"

	"github.com/quasilyte/go-ruleguard/ruleguard"
)

// Version contains extra version info.
// It's initialized via ldflags -X when ruleguard is built with Make.
// Can contain a git hash (dev builds) or a version tag (release builds).
var Version string

func docString() string { _ = "STUB: not implemented"; return "" }

// Analyzer exports ruleguard as an analysis-compatible object.
var Analyzer = &analysis.Analyzer{
	Name: "ruleguard",
	Doc:  docString(),
	Run:  runAnalyzer,
}

// ForceNewEngine disables engine cache optimization.
// This should only be useful for analyzer testing.
var ForceNewEngine = false

var runnerStatePool sync.Pool

var (
	globalEngineMu      sync.Mutex
	globalEngine        *ruleguard.Engine
	globalEngineErrored bool
)

var (
	flagRules   string
	flagE       string
	flagEnable  string
	flagDisable string

	flagGoVersion string

	flagDebug              string
	flagDebugFunc          string
	flagDebugImports       bool
	flagDebugEnableDisable bool
)

func init() {
	Analyzer.Flags.StringVar(&flagDebugFunc, "debug-func", "", "[experimental!] enable debug for the specified bytecode function")
	Analyzer.Flags.StringVar(&flagDebug, "debug-group", "", "[experimental!] enable debug for the specified matcher function")
	Analyzer.Flags.BoolVar(&flagDebugImports, "debug-imports", false, "[experimental!] enable debug for rules compile-time package lookups")
	Analyzer.Flags.BoolVar(&flagDebugEnableDisable, "debug-enable-disable", false, "[experimental!] enable debug for -enable/-disable related info")

	Analyzer.Flags.StringVar(&flagGoVersion, "go", "", "select the Go version to target; leave as string for the latest")

	Analyzer.Flags.StringVar(&flagRules, "rules", "", "comma-separated list of ruleguard file paths")
	Analyzer.Flags.StringVar(&flagE, "e", "", "execute a single rule from a given string")
	Analyzer.Flags.StringVar(&flagEnable, "enable", "<all>", "comma-separated list of enabled groups or '<all>' to enable everything")
	Analyzer.Flags.StringVar(&flagDisable, "disable", "", "comma-separated list of groups to be disabled")
}

func debugPrint(s string) { _ = "STUB: not implemented"; return }

func runAnalyzer(pass *analysis.Pass) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This condition will trigger only if we failed to init
// the engine. Return without an error as other analysis
// pass probably reported init error by this moment.

func prepareEngine() (*ruleguard.Engine, error) { _ = "STUB: not implemented"; return nil, nil }

// If we already failed once, don't try again to avoid #167.

func newEngine() (*ruleguard.Engine, error) { _ = "STUB: not implemented"; return nil, nil }
