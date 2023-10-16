package operation

//go:generate stringer -type=Operation
type Operation int

const (
	Unknown Operation = iota
	Add
	Sub
	Mul
	Div
	Pow
	Negate
	LogicNegate
)

func (o Operation) AsUnary() Operation {
	switch o {
	case Sub:
		return Negate
	case LogicNegate:
		return LogicNegate
	}

	return Unknown
}
