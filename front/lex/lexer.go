package lex

import (
	"nur/front/keyword"
	"nur/internal/comperr"
	"nur/internal/position"
)

type Lexer struct {
	stream   string
	chars    *treeNode
	pos      position.Position
	previous Lexeme
}

func NewLexer(stream string) *Lexer {
	return &Lexer{
		stream: stream,
		chars:  newCharsTree(),
	}
}

func (l *Lexer) Lex() (lexemes []Lexeme, err *comperr.Error) {
	for len(l.stream) > 0 {
		l.pos.Begin = l.pos.End

		switch char := l.stream[0]; {
		case char == ' ' || char == '\t' || char == '\r':
			l.pos.End++
			l.trimLeft(1)
		case char == '\n':
			l.pos.Line++
			l.pos.Begin = 0
			l.pos.End = 0
			l.trimLeft(1)
		case isDigit(char):
			lexemes = append(lexemes, l.parseInt())
		case isLetter(char) || char == '_':
			lexemes = append(lexemes, l.parseId())
		case l.chars.HasPrefix(string(l.stream[0])):
			lexeme := l.parseChar()
			if lexeme.Type == Unknown {
				return nil, &comperr.Error{
					Message:  "unknown operator",
					Position: lexeme.Position,
				}
			}

			lexemes = append(lexemes, lexeme)
		default:
			l.pos.End++

			return nil, &comperr.Error{
				Message:  "unknown lexeme",
				Position: l.pos,
			}
		}
	}

	return lexemes, nil
}

func (l *Lexer) trimLeft(n int) string {
	prefix := l.stream[:n]
	l.stream = l.stream[n:]

	return prefix
}

func (l *Lexer) parseInt() Lexeme {
	for i := 0; i < len(l.stream); i++ {
		if !isDigit(l.stream[i]) {
			return Lexeme{
				Type:     Int,
				Value:    l.trimLeft(i),
				Position: l.pos,
			}
		}

		l.pos.End++
	}

	return Lexeme{
		Type:     Int,
		Value:    l.trimLeft(len(l.stream)),
		Position: l.pos,
	}
}

func (l *Lexer) parseId() Lexeme {
	for i := 0; i < len(l.stream); i++ {
		switch {
		case isDigit(l.stream[i]) || isLetter(l.stream[i]) || l.stream[i] == '_':
		default:
			id := l.trimLeft(i)
			if keyword.Keywords[id] {
				return Lexeme{
					Type:     Keyword,
					Value:    id,
					Position: l.pos,
				}
			}

			return Lexeme{
				Type:     Id,
				Value:    id,
				Position: l.pos,
			}
		}

		l.pos.End++
	}

	return Lexeme{
		Type:     Id,
		Value:    l.trimLeft(len(l.stream)),
		Position: l.pos,
	}
}

func (l *Lexer) parseChar() Lexeme {
	for i := 1; i <= len(l.stream); i++ {
		if !l.chars.HasPrefix(l.stream[:i]) {
			return Lexeme{
				Type:     l.chars.Match(l.stream[:i-1]),
				Value:    l.trimLeft(i - 1),
				Position: l.pos,
			}
		}

		l.pos.End++
	}

	if _, exists := chars[l.stream]; !exists {
		return Lexeme{
			Type:     Unknown,
			Value:    l.trimLeft(len(l.stream)),
			Position: l.pos,
		}
	}

	return Lexeme{
		Type:     l.chars.Match(l.stream),
		Value:    l.trimLeft(len(l.stream)),
		Position: l.pos,
	}
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}
