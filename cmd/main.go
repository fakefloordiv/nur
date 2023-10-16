package main

import (
	"fmt"
	"nur/cmd/calculator"
	"nur/front/lex"
	"nur/front/parser"
	"nur/internal/comperr"
	"strings"
)

func main() {
	const code = `
fn sum(a int, b int) int {
	var c = 10
	return a + b + c
	c = 20
}

fn main() int {
	var a = -5
	return sum(a, 6)
}
`
	lexemes, err := lex.NewLexer(code).Lex()
	if err != nil {
		printError(code, err)
		return
	}

	fmt.Println("lexemes:", lexemes)

	ast, err := parser.NewParser(lexemes).Parse()
	if err != nil {
		printError(code, err)
		return
	}

	interpreter := calculator.NewInterpreter()
	fmt.Println(interpreter.Execute(ast))
}

func printError(code string, err *comperr.Error) {
	lines := strings.Split(code, "\n")
	fmt.Println(lines[err.Line])

	var tabs int
	for _, char := range lines[err.Line] {
		if char == '\t' {
			tabs++
		}
	}

	idents := strings.Repeat(" ", err.Begin-tabs) + strings.Repeat("\t", tabs)
	fmt.Printf("%s%s\n", idents, strings.Repeat("^", err.End-err.Begin))
	fmt.Printf("<code>:%d:%d-%d: %s\n", err.Line+1, err.Begin, err.End, err.Message)
}
