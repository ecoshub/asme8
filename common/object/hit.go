package object

import (
	"bytes"
	"encoding/binary"
)

type Position struct {
	symbol         string
	offset         uint16
	optionalOffset uint16
	missing        bool
	size           uint8
}

// GetSymbol returns the symbol of the position.
func (p *Position) GetSymbol() string {
	return p.symbol
}

// GetOffset returns the offset of the position.
func (p *Position) GetOffset() uint16 {
	return p.offset
}

// GetOptionalOffset returns the optional offset of the position.
func (p *Position) GetOptionalOffset() uint16 {
	return p.optionalOffset
}

// IsMissing returns whether the position is missing.
func (p *Position) IsMissing() bool {
	return p.missing
}

// GetSize returns the size of the position.
func (p *Position) GetSize() uint8 {
	return p.size
}

func (h *Position) Encode(buf *bytes.Buffer) error {
	if err := binary.Write(buf, binary.LittleEndian, uint16(len(h.symbol))); err != nil {
		return err
	}
	if _, err := buf.WriteString(h.symbol); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, h.offset); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, h.optionalOffset); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, h.missing); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, h.size); err != nil {
		return err
	}
	return nil
}

func (h *Position) Decode(buf *bytes.Reader) error {
	var symbolLen uint16
	if err := binary.Read(buf, binary.LittleEndian, &symbolLen); err != nil {
		return err
	}
	symbol := make([]byte, symbolLen)
	if _, err := buf.Read(symbol); err != nil {
		return err
	}
	h.symbol = string(symbol)
	if err := binary.Read(buf, binary.LittleEndian, &h.offset); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &h.optionalOffset); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &h.missing); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &h.size); err != nil {
		return err
	}
	return nil
}
