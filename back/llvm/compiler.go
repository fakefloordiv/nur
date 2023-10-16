package llvm

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

type Compiler struct {
	funcs map[string]ir.Func
	vars  map[string]value.Value
}

func (c *Compiler) Compile() {

}
