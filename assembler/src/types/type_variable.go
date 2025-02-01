package types

type Variable struct {
	Name          string
	Val           *Value
	Offset        int
	Unresolved    bool
	ReferenceName string
}

func NewVariable(name string, val *Value) *Variable {
	return &Variable{Name: name, Val: val}
}
