package quasigo

const maxFuncLocals = 8

// pop2 removes the two top stack elements and returns them.
//
// Note that it returns the popped elements in the reverse order
// to make it easier to map the order in which they were pushed.
func (s *ValueStack) pop2() (second, top interface{}) { _ = "STUB: not implemented"; return nil, nil }

func (s *ValueStack) popInt2() (second, top int) { _ = "STUB: not implemented"; return 0, 0 }

// top returns top of the stack without popping it.
func (s *ValueStack) top() interface{} { _ = "STUB: not implemented"; return nil }

func (s *ValueStack) topInt() int { _ = "STUB: not implemented"; return 0 }

// dup copies the top stack element.
// Identical to s.Push(s.Top()), but more concise.
func (s *ValueStack) dup() { _ = "STUB: not implemented"; return }

// discard drops the top stack element.
// Identical to s.Pop() without using the result.
func (s *ValueStack) discard() { _ = "STUB: not implemented"; return }

func eval(env *EvalEnv, fn *Func, top, intTop int) CallResult {
	_ = "STUB: not implemented"
	return *new(CallResult)
}
