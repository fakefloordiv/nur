package llcode

import (
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"nur/front/operation"
)

type Block any

type IntLit struct {
	Value int64
}

type Id struct {
	Value string
}

type FnDef struct {
	Name    string
	Args    []FnDefArg
	Returns types.Type
}

type FnDefArg struct {
	Name string
	Type types.Type
}

type FnCall struct {
	Target Block
	Args   []Block
}

type VarDef struct {
	Name  string
	Type  types.Type
	Value Block
}

type VarAssign struct {
	Name  string
	Value Block
}

type Operation struct {
	Left, Right value.Value
	Op          operation.Operation
}
