package object

import (
	"bytes"
	"encoding/binary"
)

type Hit struct {
	symbol         string
	offset         uint16
	optionalOffset uint16
	missing        bool
	size           uint8
}

func (h *Hit) Encode(buf *bytes.Buffer) error {
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

func (h *Hit) Decode(buf *bytes.Reader) error {
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
