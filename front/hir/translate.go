package hir

import (
	"nur/front/parser/ast"
	"nur/internal/idgen"
)

type Translate struct {
	gen idgen.Generator
}

func (t *Translate) Translate(program ast.Program) (p Program) {
	for _, node := range program {
		p = append(p, translateNode(node))
	}

	return p
}

func translateNode(node ast.Node) Node {
	switch node.(type) {
	case ast.Int:
		return IntLit{Value: node.(ast.Int).Value}
	case ast.Id:
		return Id{Value: node.(ast.Id).Name}
	case ast.BinOp:
		binOp := node.(ast.BinOp)

		return BinOp{
			Op:    0,
			Left:  translateNode(binOp.Left),
			Right: translateNode(binOp.Right),
		}
	case ast.UnOp:
		unOp := node.(ast.UnOp)

		return UnOp{
			Op:    unOp.Op,
			Right: unOp.Value,
		}
	case ast.FnDef:

	}
}

func translateFnDef(fnDef ast.FnDef) Func {

}
