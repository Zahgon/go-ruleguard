package goenv

func Read() (map[string]string, error) {
	_ = "STUB: not implemented"
	// pass in a fixed set of var names to avoid needing to unescape output
	// pass in literals here instead of a variable list to avoid security linter warnings about command injection
	return nil, nil
}

func parseGoEnv(varNames []string, data []byte) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
