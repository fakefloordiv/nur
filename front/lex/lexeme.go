package lex

import (
	"fmt"
	"nur/front/operation"
	"nur/internal/position"
)

type Lexeme struct {
	Type
	Value string
	position.Position
}

func (l Lexeme) String() string {
	switch l.Type {
	case LParen, RParen, LBrace, RBrace, LQBrace, RQBrace:
		return fmt.Sprintf("(%s)", l.Type)
	}

	return fmt.Sprintf("(%s %s)", l.Type, l.Value)
}

type Type int

func (t Type) AsOperation() operation.Operation {
	switch t {
	case OpPlus:
		return operation.Add
	case OpMinus:
		return operation.Sub
	case OpSlash:
		return operation.Div
	case OpStar:
		return operation.Mul
	case OpStarStar:
		return operation.Pow
	case UnExclaim:
		return operation.LogicNegate
	}

	return operation.Unknown
}

func (t Type) IsOperator() bool {
	return t.AsOperation() != operation.Unknown
}

func (t Type) IsBrace() bool {
	switch t {
	case LParen, LBrace, LQBrace, RParen, RBrace, RQBrace:
		return true
	}

	return false
}

//go:generate stringer -type=Type

const (
	Unknown Type = iota
	Keyword
	Int
	Id
	OpPlus
	OpMinus
	OpSlash
	OpStar
	OpStarStar
	UnExclaim
	ChEq
	ChDot
	LParen
	RParen
	LBrace
	RBrace
	LQBrace
	RQBrace
	Comma
)

var chars = map[string]Type{
	"+":  OpPlus,
	"-":  OpMinus,
	"/":  OpSlash,
	"*":  OpStar,
	"**": OpStarStar,
	"!":  UnExclaim,
	"=":  ChEq,
	".":  ChDot,
	"(":  LParen,
	")":  RParen,
	"{":  LBrace,
	"}":  RBrace,
	"[":  LQBrace,
	"]":  RQBrace,
	",":  Comma,
}
