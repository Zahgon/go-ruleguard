package quasigo

import (
	"go/ast"
	"go/types"
)

func pickOp(cond bool, ifTrue, otherwise opcode) opcode {
	_ = "STUB: not implemented"
	return *new(opcode)
}

func put16(code []byte, pos, value int) { _ = "STUB: not implemented"; return }

func decode16(code []byte, pos int) int { _ = "STUB: not implemented"; return 0 }

func typeIsInt(typ types.Type) bool { _ = "STUB: not implemented"; return false }

func typeIsString(typ types.Type) bool { _ = "STUB: not implemented"; return false }

func walkBytecode(code []byte, fn func(pc int, op opcode)) { _ = "STUB: not implemented"; return }

func identName(n ast.Expr) string { _ = "STUB: not implemented"; return "" }
