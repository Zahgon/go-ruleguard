//go:build ignore
// +build ignore

package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

// Do not log errors as unstructured fields.
// Before:
//
//	logger.Infof("Unable to create profile. Error: %v", err)
//
// After:
//
//	logger.Info("Unable to create profile", zap.Error(err))
func logWithUnstructuredError(m dsl.Matcher) { _ = "STUB: not implemented"; return }
