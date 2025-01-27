package object

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"sort"
)

type Tracker struct {
	segment   string
	defined   map[string]*Symbol
	positions map[uint16]*Position
	code      string
	bin       []uint8
}

func NewTracker() *Tracker {
	return &Tracker{
		defined:   make(map[string]*Symbol),
		positions: make(map[uint16]*Position),
	}
}

func (t *Tracker) GetSymbols() map[string]*Symbol {
	return t.defined
}

func (t *Tracker) GetPositions() map[uint16]*Position {
	return t.positions
}

func (t *Tracker) GetSegment() string {
	return t.segment
}

func (t *Tracker) GetCode() string {
	return t.code
}

func (t *Tracker) GetBin() []uint8 {
	return t.bin
}

func (t *Tracker) AttachCode(code string) {
	t.code = code
}

func (t *Tracker) AttachBin(bin []byte) {
	t.bin = bin
}

func (t *Tracker) SetSegment(segment string) {
	t.segment = segment
}

func (t *Tracker) DefineSymbol(symbol string) {
	t.defined[symbol] = NewSymbol(symbol)
}

func (t *Tracker) SymbolHit(symbol string, offset uint16, size uint8, optionalOffset uint16) {
	t.positions[offset] = &Position{symbol: symbol, offset: offset, size: size, optionalOffset: optionalOffset}
}

func (t *Tracker) SetMissing(symbol string, offset uint16, size uint8, optionalOffset uint16) {
	t.positions[offset] = &Position{symbol: symbol, offset: offset, size: size, missing: true, optionalOffset: optionalOffset}
}

func (t *Tracker) SetIndex(symbol string, index uint16) {
	s := t.GetOrDefineSymbol(symbol)
	s.SetIndex(index)
}

func (t *Tracker) SetType(symbol string, _type uint8) {
	s := t.GetOrDefineSymbol(symbol)
	s.SetType(_type)
}

func (t *Tracker) SetExtern(symbol string) {
	s := t.GetOrDefineSymbol(symbol)
	s.share = SYMBOL_SHARE_STATUS_EXTERN
}

func (t *Tracker) SetGlobal(symbol string) {
	s := t.GetOrDefineSymbol(symbol)
	s.share = SYMBOL_SHARE_STATUS_GLOBAL
}

func (t *Tracker) GetOrDefineSymbol(symbol string) *Symbol {
	s, exists := t.defined[symbol]
	if !exists {
		s = NewSymbol(symbol)
		t.defined[symbol] = s
	}
	return s
}

func (t *Tracker) GetDefinedSymbol(symbol string) (*Symbol, bool) {
	s, exists := t.defined[symbol]
	return s, exists
}

func (t *Tracker) GetSymbolSize(symbol string, def uint8) uint8 {
	for _, e := range t.positions {
		if e.symbol == symbol {
			if e.size != 0 {
				return e.size
			}
		}
	}
	return def
}

func (t *Tracker) Print() {
	fmt.Println("FILE SEGMENT:")
	fmt.Println(t.segment)
	fmt.Println()

	sortedSymbols := SortSymbolMap(t.defined)
	fmt.Println("SYMBOL TABLE:")
	fmt.Println("OFFSET      ACCESS     TYPE     SYMBOL")
	for _, key := range sortedSymbols {
		d := t.defined[key]
		fmt.Printf("%04x        %03b        %d        <%s>\n", d.index, d.share, d._type, d.symbol)
	}
	fmt.Println()

	sortedHits := SortHitMap(t.positions)
	fmt.Println("POSITIONS:")
	fmt.Println("OFFSET     SIZE     EXTRA   missing    SYMBOL ")
	for _, key := range sortedHits {
		e := t.positions[key]
		fmt.Printf("%04x       %2d       %04x    %-5v      <%s>\n", e.offset, e.size, e.optionalOffset, e.missing, e.symbol)
	}
	fmt.Println()

	fmt.Println(t.code)
}

func SortSymbolMap(m map[string]*Symbol) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]].index < m[keys[j]].index
	})

	return keys
}

func SortHitMap(m map[uint16]*Position) []uint16 {
	keys := make([]uint16, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]].offset < m[keys[j]].offset
	})

	return keys
}

func (t *Tracker) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Encode segment
	err := binary.Write(buf, binary.LittleEndian, uint32(len(t.segment)))
	if err != nil {
		return nil, err
	}
	_, err = buf.WriteString(t.segment)
	if err != nil {
		return nil, err
	}

	// Encode defined symbols
	err = binary.Write(buf, binary.LittleEndian, uint32(len(t.defined)))
	if err != nil {
		return nil, err
	}
	for key, symbol := range t.defined {
		err = binary.Write(buf, binary.LittleEndian, uint32(len(key)))
		if err != nil {
			return nil, err
		}
		_, err = buf.WriteString(key)
		if err != nil {
			return nil, err
		}
		err = symbol.Encode(buf)
		if err != nil {
			return nil, err
		}
	}

	// Encode positions
	err = binary.Write(buf, binary.LittleEndian, uint32(len(t.positions)))
	if err != nil {
		return nil, err
	}
	for key, hit := range t.positions {
		err = binary.Write(buf, binary.LittleEndian, key)
		if err != nil {
			return nil, err
		}
		err = hit.Encode(buf)
		if err != nil {
			return nil, err
		}
	}

	// Encode code
	err = binary.Write(buf, binary.LittleEndian, uint32(len(t.code)))
	if err != nil {
		return nil, err
	}
	_, err = buf.WriteString(t.code)
	if err != nil {
		return nil, err
	}

	// Encode bin
	err = binary.Write(buf, binary.LittleEndian, uint32(len(t.bin)))
	if err != nil {
		return nil, err
	}
	_, err = buf.Write(t.bin)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (t *Tracker) Decode(data []byte) error {
	buf := bytes.NewReader(data)

	// Decode segment
	var segmentLen uint32
	err := binary.Read(buf, binary.LittleEndian, &segmentLen)
	if err != nil {
		return err
	}
	segment := make([]byte, segmentLen)
	_, err = buf.Read(segment)
	if err != nil {
		return err
	}
	t.segment = string(segment)

	// Decode defined symbols
	var definedLen uint32
	err = binary.Read(buf, binary.LittleEndian, &definedLen)
	if err != nil {
		return err
	}
	t.defined = make(map[string]*Symbol, definedLen)
	for i := uint32(0); i < definedLen; i++ {
		var keyLen uint32
		err = binary.Read(buf, binary.LittleEndian, &keyLen)
		if err != nil {
			return err
		}
		key := make([]byte, keyLen)
		_, err = buf.Read(key)
		if err != nil {
			return err
		}
		symbol := new(Symbol)
		err = symbol.Decode(buf)
		if err != nil {
			return err
		}
		t.defined[string(key)] = symbol
	}

	// Decode positions
	var positionsLen uint32
	err = binary.Read(buf, binary.LittleEndian, &positionsLen)
	if err != nil {
		return err
	}
	t.positions = make(map[uint16]*Position, positionsLen)
	for i := uint32(0); i < positionsLen; i++ {
		var key uint16
		err = binary.Read(buf, binary.LittleEndian, &key)
		if err != nil {
			return err
		}
		hit := new(Position)
		err = hit.Decode(buf)
		if err != nil {
			return err
		}
		t.positions[key] = hit
	}

	// Decode code
	var codeLen uint32
	err = binary.Read(buf, binary.LittleEndian, &codeLen)
	if err != nil {
		return err
	}
	code := make([]byte, codeLen)
	_, err = buf.Read(code)
	if err != nil {
		return err
	}
	t.code = string(code)

	// Decode bin
	var binLen uint32
	err = binary.Read(buf, binary.LittleEndian, &binLen)
	if err != nil {
		return err
	}

	if binLen == 0 {
		return nil
	}

	t.bin = make([]uint8, binLen)
	_, err = buf.Read(t.bin)
	if err != nil {
		return err
	}

	return nil
}
