package lex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLexer_Lex(t *testing.T) {
	tcs := []struct {
		Name   string
		Sample string
		Want   []Lexeme
	}{
		{
			Name:   "simple expression",
			Sample: "1+1",
			Want: []Lexeme{
				{Type: Int, Value: "1"}, {Type: OpPlus}, {Type: Int, Value: "1"},
			},
		},
		{
			Name:   "composite operator",
			Sample: "1**1",
			Want: []Lexeme{
				{Type: Int, Value: "1"}, {Type: OpStarStar}, {Type: Int, Value: "1"},
			},
		},
		{
			Name:   "single identifier",
			Sample: "hello",
			Want: []Lexeme{
				{Type: Id, Value: "hello"},
			},
		},
		{
			Name:   "single integer",
			Sample: "5",
			Want: []Lexeme{
				{Type: Int, Value: "5"},
			},
		},
		{
			Name:   "nested braces",
			Sample: "{[()]}",
			Want: []Lexeme{
				{Type: LBrace}, {Type: LQBrace}, {Type: LParen}, {Type: RParen}, {Type: RQBrace}, {Type: RBrace},
			},
		},
		{
			Name:   "composed expression",
			Sample: "1 ** (hello - 51)",
			Want: []Lexeme{
				{Type: Int, Value: "1"}, {Type: OpStarStar}, {Type: LParen}, {Type: Id, Value: "hello"},
				{Type: OpMinus}, {Type: Int, Value: "51"}, {Type: RParen},
			},
		},
	}

	for _, tc := range tcs {
		lexemes, err := NewLexer(tc.Sample).Lex()
		require.Nil(t, err, tc.Name)
		wantSequence(t, tc.Name, lexemes, tc.Want)
	}
}

func wantSequence(t *testing.T, name string, given, want []Lexeme) {
	require.Equalf(t, len(given), len(want), "%s: lengths of wanted and given lexemes are mismatching", name)

	for i, lexeme := range given {
		require.Equalf(t, want[i].Type, lexeme.Type, "%s: mismatched: %s and %s", name, want[i].Type, lexeme.Type)
		if len(want[i].Value) > 0 {
			require.Equal(t, want[i].Value, lexeme.Value, name)
		}
	}
}
