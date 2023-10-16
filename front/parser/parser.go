package parser

import (
	"fmt"
	"nur/front/keyword"
	"nur/front/lex"
	"nur/front/operation"
	"nur/front/parser/ast"
	"nur/internal/comperr"
	"nur/internal/position"
	"nur/internal/walk"
	"strconv"
)

type Parser struct {
	stream *walk.Walker[lex.Lexeme]
}

func NewParser(lexemes []lex.Lexeme) *Parser {
	return &Parser{
		stream: walk.NewWalker(lexemes),
	}
}

func (p *Parser) Parse() (program ast.Program, err *comperr.Error) {
	var nodes []ast.Node

	for p.stream.Current().Type != lex.Unknown {
		node, err := p.stmt()
		if err != nil {
			return nil, err
		}

		nodes = append(nodes, node)
	}

	return nodes, nil
}

func (p *Parser) stmt() (ast.Node, *comperr.Error) {
	switch lexeme := p.stream.Current(); lexeme.Type {
	case lex.Keyword:
		switch lexeme.Value {
		case keyword.Fn:
			p.stream.Next()
			var fn ast.FnDef

			name, err := p.match(lex.Id)
			if err != nil {
				return nil, err
			}

			fn.Name = name.Value
			fn.Pos.Line = name.Line
			fn.Pos.Begin = name.Begin

			if _, err = p.match(lex.LParen); err != nil {
				return nil, err
			}

			for p.stream.Current().Type != lex.RParen {
				argName, err := p.match(lex.Id)
				if err != nil {
					return nil, err
				}

				argType, err := p.match(lex.Id)
				if err != nil {
					return nil, err
				}

				fn.Args = append(fn.Args, ast.NameType{
					Name: argName.Value,
					Type: argType.Value,
					Pos: position.Position{
						Line:  argName.Position.Line,
						Begin: argName.Position.Begin,
						End:   argType.Position.End,
					},
				})

				// optionally skip comma. If there's no comma, stream won't be advanced
				_, _ = p.match(lex.Comma)
			}

			rparen, err := p.match(lex.RParen)
			if err != nil {
				return nil, err
			}

			fn.Pos.End = rparen.End

			if p.stream.Current().Type == lex.Id {
				ret, _ := p.match(lex.Id)
				fn.Returns.Type = ret.Value
				fn.Pos.End = ret.End
			}

			body, err := p.fnBody()
			fn.Body = body

			return fn, err
		default:
			return nil, &comperr.Error{
				Message:  "unexpected statement",
				Position: lexeme.Position,
			}
		}
	}

	return nil, &comperr.Error{
		Message:  "unexpected expression",
		Position: p.stream.Current().Position,
	}
}

func (p *Parser) fnBody() (ast.Node, *comperr.Error) {
	switch lexeme := p.stream.Current(); lexeme.Type {
	case lex.LBrace:
		p.stream.Next()
		var block []ast.Node

		for p.stream.Current().Type != lex.RBrace {
			expr, err := p.expr()
			if err != nil {
				return nil, err
			}

			block = append(block, expr)
		}

		_, _ = p.match(lex.RBrace)

		return ast.Block{Exprs: block, Position: lexeme.Position}, nil
	default:
		return p.expr()
	}
}

func (p *Parser) expr() (ast.Node, *comperr.Error) {
	node, err := p.term()
	if err != nil {
		return nil, err
	}

	for {
		switch lexeme := p.stream.Current(); lexeme.Type {
		case lex.OpPlus, lex.OpMinus:
			p.stream.Next()
			operator := lexeme.AsOperation()
			right, err := p.expr()
			if err != nil {
				return nil, err
			}

			node = ast.BinOp{
				Op:    operator,
				Left:  node,
				Right: right,
			}
		case lex.ChEq:
			p.stream.Next()
			rhs, err := p.expr()

			return ast.VarAssign{
				Lhs: node,
				Rhs: rhs,
			}, err
		default:
			return node, nil
		}
	}
}

func (p *Parser) term() (ast.Node, *comperr.Error) {
	node, err := p.power()
	if err != nil {
		return nil, err
	}

	for isOneOf(p.stream.Current().Type, lex.OpStar, lex.OpSlash) {
		operator := p.stream.Current().Type.AsOperation()
		p.stream.Next()
		right, err := p.term()
		if err != nil {
			return nil, err
		}

		node = ast.BinOp{
			Op:    operator,
			Left:  node,
			Right: right,
		}
	}

	return node, nil
}

func (p *Parser) power() (ast.Node, *comperr.Error) {
	node, err := p.factor()
	if err != nil {
		return nil, err
	}

	for {
		lexeme := p.stream.Current()
		switch lexeme.Type {
		case lex.OpStarStar:
			p.stream.Next()
			right, err := p.power()
			if err != nil {
				return nil, err
			}

			node = ast.BinOp{
				Op:    operation.Pow,
				Left:  node,
				Right: right,
			}
		case lex.LParen:
			p.stream.Next()
			var args []ast.Node

			for p.stream.Current().Type != lex.RParen {
				arg, err := p.expr()
				if err != nil {
					return nil, err
				}

				args = append(args, arg)
				_, _ = p.match(lex.Comma)
			}

			p.stream.Next()

			node = ast.FnCall{
				Target: node,
				Args:   args,
			}
		default:
			return node, nil
		}
	}
}

func (p *Parser) factor() (ast.Node, *comperr.Error) {
	lexeme := p.stream.Current()
	p.stream.Next()
	switch lexeme.Type {
	case lex.Int:
		num, err := strconv.Atoi(lexeme.Value)
		if err != nil {
			return nil, &comperr.Error{
				Message:  "invalid integer literal",
				Position: lexeme.Position,
			}
		}

		return ast.Int{Value: int64(num)}, nil
	case lex.Id:
		return ast.Id{Name: lexeme.Value}, nil
	case lex.LParen:
		stmt, err := p.expr()
		if err != nil {
			return nil, err
		}

		_, err = p.match(lex.RParen)

		return stmt, err
	case lex.Keyword:
		switch lexeme.Value {
		case keyword.Var:
			return p.varDef()
		case keyword.Return:
			value, err := p.expr()

			return ast.Return{
				Value: value,
			}, err
		}
	}

	if lexeme.Type.IsOperator() {
		op := lexeme.Type.AsOperation().AsUnary()
		if op == operation.Unknown {
			return nil, &comperr.Error{
				Message:  "bad unary operator",
				Position: lexeme.Position,
			}
		}

		right, err := p.power()

		return ast.UnOp{
			Op:    op,
			Value: right,
		}, err
	}

	return nil, &comperr.Error{
		Message:  "unexpected lexeme",
		Position: lexeme.Position,
	}
}

func (p *Parser) varDef() (ast.Node, *comperr.Error) {
	name, err := p.match(lex.Id)
	if err != nil {
		return nil, &comperr.Error{
			Message:  "must be identifier",
			Position: name.Position,
		}
	}

	_, err = p.match(lex.ChEq)
	if err != nil {
		return nil, err
	}

	value, err := p.expr()

	return ast.VarDef{
		Name:  name.Value,
		Value: value,
	}, err
}

func (p *Parser) match(desired lex.Type) (lex.Lexeme, *comperr.Error) {
	current := p.stream.Current()

	if current.Type != desired {
		return current, &comperr.Error{
			Message:  fmt.Sprintf("wanted %s, got %s", desired, current.Type),
			Position: current.Position,
		}
	}

	p.stream.Next()

	return current, nil
}

func isOneOf(one lex.Type, ofs ...lex.Type) bool {
	for _, of := range ofs {
		if one == of {
			return true
		}
	}

	return false
}
