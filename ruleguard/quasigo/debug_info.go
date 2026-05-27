package quasigo

type debugInfo struct {
	funcs map[*Func]funcDebugInfo
}

type funcDebugInfo struct {
	paramNames    []string
	intParamNames []string
	localNames    []string
}

func newDebugInfo() *debugInfo { _ = "STUB: not implemented"; return nil }
