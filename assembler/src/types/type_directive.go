package types

type Directive struct {
	Raw      string
	Type     string
	Position uint16
	Offset   uint16
	Single   bool
	Values   []*Value
}
