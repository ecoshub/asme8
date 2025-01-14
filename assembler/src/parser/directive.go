package parser

type Directive struct {
	raw      string
	code     string
	position uint64
	offset   uint64
	single   bool
	inm      []uint16
}
