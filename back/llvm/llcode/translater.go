package llcode

import (
	"nur/front/parser/ast"
	"nur/internal/comperr"
)

func Translate(program ast.Program) (funcs []FnDef, err *comperr.Error) {
	for _, node := range program {
		switch node.(type) {
		case ast.FnDef:
			fndef := node.(ast.FnDef)
			var args []FnDefArg

			for _, arg := range fndef.Args {

			}

		}
	}
}
