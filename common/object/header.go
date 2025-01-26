package object

import (
	"bytes"
	"encoding/binary"
)

const (
	MAGIC_NUMBER_U32 uint32 = 2135247942
)

type Header struct {
	// 4 bytes
	magic uint32
	// 2 bytes
	version uint16
}

func NewStdHeader(version uint16) *Header {
	return &Header{
		magic:   MAGIC_NUMBER_U32,
		version: version,
	}
}

func (h *Header) Decode(buf *bytes.Reader) error {
	// Decode magic
	if err := binary.Read(buf, binary.BigEndian, &h.magic); err != nil {
		return err
	}

	// Decode version
	if err := binary.Read(buf, binary.LittleEndian, &h.version); err != nil {
		return err
	}

	return nil
}

func (h *Header) Encode(buf *bytes.Buffer) error {
	// Encode magic
	if err := binary.Write(buf, binary.BigEndian, h.magic); err != nil {
		return err
	}

	// Encode version
	if err := binary.Write(buf, binary.LittleEndian, h.version); err != nil {
		return err
	}

	return nil
}
