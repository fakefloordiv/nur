package llcode

import "github.com/llir/llvm/ir/types"

var typesMapping = map[string]types.Type{
	"bool":  types.I1,
	"int":   types.I64,
	"int8":  types.I8,
	"int32": types.I32,
	"int64": types.I64,
}
