package ruleguard

import (
	"go/types"

	"github.com/quasilyte/go-ruleguard/ruleguard/ir"
)

func convertAST(ctx *LoadContext, imp *goImporter, filename string, src []byte) (*ir.File, *types.Package, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
