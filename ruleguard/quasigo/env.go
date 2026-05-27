package quasigo

type funcKey struct {
	qualifier string
	name      string
}

func (k funcKey) String() string { _ = "STUB: not implemented"; return "" }

type nativeFunc struct {
	mappedFunc func(*ValueStack)
	name       string // Needed for the readable disasm
}

func newEnv() *Env { _ = "STUB: not implemented"; return nil }

func (env *Env) addNativeFunc(key funcKey, f func(*ValueStack)) { _ = "STUB: not implemented"; return }

func (env *Env) addFunc(key funcKey, f *Func) { _ = "STUB: not implemented"; return }
