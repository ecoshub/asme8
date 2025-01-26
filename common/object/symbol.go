package object

import (
	"bytes"
	"encoding/binary"
)

const (
	SYMBOL_SHARE_STATUS_LOCAL  uint8 = 1
	SYMBOL_SHARE_STATUS_EXTERN uint8 = 1 << 1
	SYMBOL_SHARE_STATUS_GLOBAL uint8 = 1 << 2
)

const (
	SYMBOL_RESOLVE_STATUS_NOT_RESOLVED uint8 = 0
	SYMBOL_RESOLVE_STATUS_RESOLVED     uint8 = 1
)

type Symbol struct {
	symbol string
	index  uint16
	share  uint8
}

func NewSymbol(symbol string) *Symbol {
	return &Symbol{
		symbol: symbol,
		share:  SYMBOL_SHARE_STATUS_LOCAL,
	}
}

func (s *Symbol) SetShare(status uint8) {
	s.share = status
}

func (s *Symbol) SetIndex(index uint16) {
	s.index = index
}

func (s *Symbol) Encode(buf *bytes.Buffer) error {
	// Encode symbol string length and value
	if err := binary.Write(buf, binary.LittleEndian, uint16(len(s.symbol))); err != nil {
		return err
	}
	if _, err := buf.WriteString(s.symbol); err != nil {
		return err
	}

	// Encode index
	if err := binary.Write(buf, binary.LittleEndian, s.index); err != nil {
		return err
	}

	// Encode share
	if err := binary.Write(buf, binary.LittleEndian, s.share); err != nil {
		return err
	}

	return nil
}

func (s *Symbol) Decode(buf *bytes.Reader) error {
	// Decode symbol string length and value
	var symbolLen uint16
	if err := binary.Read(buf, binary.LittleEndian, &symbolLen); err != nil {
		return err
	}
	symbol := make([]byte, symbolLen)
	if _, err := buf.Read(symbol); err != nil {
		return err
	}
	s.symbol = string(symbol)

	// Decode index
	if err := binary.Read(buf, binary.LittleEndian, &s.index); err != nil {
		return err
	}

	// Decode share
	if err := binary.Read(buf, binary.LittleEndian, &s.share); err != nil {
		return err
	}

	return nil
}
