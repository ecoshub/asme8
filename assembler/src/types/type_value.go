package types

import (
	"strconv"
	"strings"
)

type Value struct {
	val      uint16
	size     int
	lowByte  uint8
	highByte uint8
}

func (i *Value) GetValue() uint16 {
	return i.val
}

func (i *Value) GetSize() int {
	return i.size
}

func (i *Value) GetLowByte() uint8 {
	return uint8(i.val)
}

func (i *Value) GetHighByte() uint8 {
	return uint8(i.val >> 8)
}

func (i *Value) Copy() *Value {
	return &Value{
		val:      i.val,
		size:     i.size,
		lowByte:  i.lowByte,
		highByte: i.highByte,
	}
}

func NewValue(val int64) *Value {
	im := &Value{
		val:      uint16(val),
		lowByte:  uint8(val),
		highByte: uint8(val >> 8),
	}
	if val > 0xff {
		im.size = 16
	} else {
		im.size = 8
	}
	return im
}

func ParseValue(text string) *Value {
	if len(text) == 3 {
		if text[0] == '\'' && text[2] == '\'' {
			return NewValue(int64(text[1]))
		}

		if text[0] == '"' && text[2] == '"' {
			return NewValue(int64(text[1]))
		}
	}

	if strings.HasPrefix(text, "0x") {
		text = strings.TrimPrefix(text, "0x")
		val, _ := strconv.ParseInt(text, 16, 64)
		return NewValue(val)
	}

	if strings.HasPrefix(text, "0b") {
		text = strings.TrimPrefix(text, "0b")
		val, _ := strconv.ParseInt(text, 2, 64)
		return NewValue(val)
	}

	val, _ := strconv.ParseInt(text, 10, 64)
	return NewValue(val)
}
