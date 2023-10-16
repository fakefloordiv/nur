package keyword

const (
	Fn     = "fn"
	Var    = "var"
	Return = "return"
)

var Keywords = map[string]bool{
	Fn:     true,
	Var:    true,
	Return: true,
}
