package types

type Label struct {
	Text   string
	Line   int
	Column int
	Offset uint16
}

func NewLabel(text string, line, column int, offset uint16) *Label {
	return &Label{
		Text:   text,
		Line:   line,
		Column: column,
		Offset: offset,
	}
}
