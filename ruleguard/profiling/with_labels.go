//go:build pproflabels
// +build pproflabels

package profiling

import (
	"context"
)

const LabelsEnabled = true

func EnterWithLabels(origContext context.Context, name string) { _ = "STUB: not implemented"; return }

func Leave(origContext context.Context) { _ = "STUB: not implemented"; return }
