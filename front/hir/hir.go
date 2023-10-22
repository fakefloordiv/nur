package hir

import "nur/front/operation"

type Node any

type Program []Node

type Block []Node

type Type = string

type Func struct {
	Name    string
	Args    []FuncArg
	Returns Type
	Body    Node
}

type FuncArg struct {
	Name string
	Type Type
}

type FCall struct {
	Target Node
	Args   []Node
}

type IntLit struct {
	Value int64
}

type StrLit struct {
	Value string
}

type Id struct {
	Value string
}

type VarDef struct {
	Name  string
	Value Node
}

type VarAssign struct {
	Name  string
	Value Node
}

type ArrayAssign struct {
	Name  string
	Index Node
	Value Node
}

type UnOp struct {
	Op    operation.Operation
	Right Node
}

type BinOp struct {
	Op          operation.Operation
	Left, Right Node
}
