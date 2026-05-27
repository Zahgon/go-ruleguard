package dsl

import (
	"github.com/quasilyte/go-ruleguard/dsl/types"
)

type DoContext struct{}

func (*DoContext) Var(varname string) *DoVar { _ = "STUB: not implemented"; return nil }

func (*DoContext) SetReport(report string) { _ = "STUB: not implemented"; return }

func (*DoContext) SetSuggest(suggest string) { _ = "STUB: not implemented"; return }

type DoVar struct{}

func (*DoVar) Text() string { _ = "STUB: not implemented"; return "" }

func (*DoVar) Type() types.Type { _ = "STUB: not implemented"; return *new(types.Type) }
