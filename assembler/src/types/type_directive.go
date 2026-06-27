package types

type DirectiveType string

const (
	DirectiveOrg         DirectiveType = ".org"
	DirectiveReserveByte DirectiveType = ".resb"
	DirectiveByte        DirectiveType = ".byte"
	DirectiveWord        DirectiveType = ".word"
	DirectiveASCII       DirectiveType = ".ascii"
	DirectiveASCIIZ      DirectiveType = ".asciiz"
)

type Directive struct {
	Raw      string
	Type     DirectiveType
	Position uint16
	Offset   uint16
	Values   []*Value
}
