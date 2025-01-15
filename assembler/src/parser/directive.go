package parser

type Directive struct {
	raw      string
	code     string
	position uint16
	offset   uint16
	single   bool
	inm      []uint16
}
