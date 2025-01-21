package types

type Variable struct {
	Name string
	Val  *Value
}

func NewVariable(name string, val *Value) *Variable {
	return &Variable{Name: name, Val: val}
}
