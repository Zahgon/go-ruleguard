package irprint

import (
	"bytes"
	"io"
	"reflect"

	"github.com/quasilyte/go-ruleguard/ruleguard/ir"
)

func File(w io.Writer, f *ir.File) { _ = "STUB: not implemented"; return }

type printer struct {
	buf bytes.Buffer
}

func (p *printer) writef(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (p *printer) printFile(f *ir.File) { _ = "STUB: not implemented"; return }

func (p *printer) printReflectElem(key string, v reflect.Value, insideList bool) {
	_ = "STUB: not implemented"
	return
}

func (p *printer) printReflectElemNoNewline(key string, v reflect.Value, insideList bool) bool {
	_ = "STUB: not implemented"
	return false
}

// There are tons of these, print them in a compact way.

func isCompactSlice(v reflect.Value) bool { _ = "STUB: not implemented"; return false }
