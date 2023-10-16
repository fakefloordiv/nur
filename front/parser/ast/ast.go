package ast

import (
	"nur/front/operation"
	"nur/internal/position"
)

type Node interface {
	Position() position.Position
}

type Program []Node

type Id struct {
	Name string
	Pos  position.Position
}

func (i Id) Position() position.Position {
	return i.Pos
}

type Int struct {
	Value int64
	Pos   position.Position
}

func (i Int) Position() position.Position {
	return i.Pos
}

type UnOp struct {
	Op    operation.Operation
	Value Node
	Pos   position.Position
}

func (u UnOp) Position() position.Position {
	return u.Pos
}

type BinOp struct {
	Op          operation.Operation
	Left, Right Node
}

func (b BinOp) Position() position.Position {
	left := b.Left.Position()
	right := b.Right.Position()

	return position.Position{
		Line:  left.Line,
		Begin: left.Begin,
		End:   right.End,
	}
}

type FnDef struct {
	Name    string
	Args    []NameType
	Returns NameType
	Pos     position.Position
	Body    Node
}

func (f FnDef) Position() position.Position {
	return f.Pos
}

type VarDef struct {
	Name    string
	Value   Node
	NamePos position.Position
}

func (v VarDef) Position() position.Position {
	return position.Position{
		Line:  v.NamePos.Line,
		Begin: v.NamePos.Begin,
		End:   v.Value.Position().End,
	}
}

type VarAssign struct {
	Lhs Node
	Rhs Node
}

func (v VarAssign) Position() position.Position {
	return position.Position{
		Line:  v.Lhs.Position().Line,
		Begin: v.Lhs.Position().Begin,
		End:   v.Rhs.Position().End,
	}
}

type NameType struct {
	Name, Type string
	Pos        position.Position
}

func (n NameType) Position() position.Position {
	return n.Pos
}

type Block struct {
	Exprs []Node
	Pos   position.Position
}

func (b Block) Position() position.Position {
	return b.Pos
}

type FnCall struct {
	Target Node
	Args   []Node
	Pos    position.Position
}

func (f FnCall) Position() position.Position {
	return f.Pos
}

type Return struct {
	Value Node
	Pos   position.Position
}

func (r Return) Position() position.Position {
	return r.Pos
}
