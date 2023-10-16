package calculator

import (
	"fmt"
	"math"
	"nur/front/operation"
	"nur/front/parser/ast"
	"reflect"
)

var consts = map[string]int{
	"pi":         4,
	"gravity":    10,
	"easter_egg": 42,
}

type Interpreter struct {
	vars  map[string]int
	funcs map[string]ast.FnDef
}

func NewInterpreter() Interpreter {
	return Interpreter{
		vars:  consts,
		funcs: make(map[string]ast.FnDef),
	}
}

func (i Interpreter) Execute(nodes []ast.Node) (results []int) {
	for _, node := range nodes {
		results = append(results, i.evaluate(node))
	}

	if main, found := i.funcs["main"]; found {
		return []int{i.evaluate(main.Body)}
	}

	return results
}

func (i Interpreter) evaluate(node ast.Node) int {
	switch node.(type) {
	case ast.FnDef:
		fndef := node.(ast.FnDef)
		i.funcs[fndef.Name] = fndef

		return 0
	case ast.Int:
		return int(node.(ast.Int).Value)
	case ast.Id:
		id := node.(ast.Id)

		if value, found := consts[id.Name]; found {
			return value
		}

		panic(fmt.Sprintf("name not found: %s", id.Name))
	case ast.BinOp:
		binop := node.(ast.BinOp)
		left := i.evaluate(binop.Left)
		right := i.evaluate(binop.Right)

		switch binop.Op {
		case operation.Add:
			return left + right
		case operation.Sub:
			return left - right
		case operation.Div:
			return left / right
		case operation.Mul:
			return left * right
		case operation.Pow:
			return int(math.Pow(float64(left), float64(right)))
		}

		panic("unknown operation")
	case ast.UnOp:
		unop := node.(ast.UnOp)
		right := i.evaluate(unop.Value)

		switch unop.Op {
		case operation.Negate:
			return -right
		case operation.LogicNegate:
			// I don't know how to convert bool to int. So...
			if right > 0 {
				return 0
			}

			return 1
		}

		panic("unknown unary operation")
	case ast.Block:
		block := node.(ast.Block)
		var result int

		for _, expr := range block.Exprs {
			if ret, ok := expr.(ast.Return); ok {
				return i.evaluate(ret.Value)
			}

			result = i.evaluate(expr)
		}

		return result
	case ast.Return:
		ret := node.(ast.Return)

		return i.evaluate(ret.Value)
	case ast.VarDef:
		varDef := node.(ast.VarDef)
		value := i.evaluate(varDef.Value)
		i.vars[varDef.Name] = value

		return 0
	case ast.FnCall:
		fncall := node.(ast.FnCall)
		name, ok := fncall.Target.(ast.Id)
		if !ok {
			panic(fmt.Sprintf("not able to perform call to %s (possibly only to identifier)", fncall.Target))
		}

		fn := i.funcs[name.Name]
		var args []int
		for _, arg := range fncall.Args {
			args = append(args, i.evaluate(arg))
		}

		if len(args) != len(fn.Args) {
			panic(fmt.Sprintf("want %d args, got %d", len(fn.Args), len(args)))
		}

		for index, argName := range fn.Args {
			i.vars[argName.Name] = args[index]
		}

		return i.evaluate(fn.Body)
	}

	panic(fmt.Sprintf("unknown node: %s", reflect.TypeOf(node)))
}
