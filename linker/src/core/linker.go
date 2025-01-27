package core

import (
	"asme8/common/object"
	"asme8/linker/src/config"
	"encoding/binary"
	"fmt"
	"os"
)

const (
	MEMORY_SIZE int = 0x2000
)

type Linker struct {
	config            *config.Config
	objects           []*object.ELF
	memory            []byte
	lastMemoryOffsets map[string]uint16
	segmentOffsets    map[string]uint16
	globals           map[string][]*object.Symbol
	externs           map[string][]*object.Symbol
	missing           map[string][]*object.Symbol
}

func NewLinker(config *config.Config, objects ...*object.ELF) *Linker {
	return &Linker{
		config:            config,
		objects:           objects,
		memory:            make([]byte, MEMORY_SIZE),
		lastMemoryOffsets: make(map[string]uint16),
		segmentOffsets:    make(map[string]uint16),
		globals:           make(map[string][]*object.Symbol),
		externs:           make(map[string][]*object.Symbol),
		missing:           make(map[string][]*object.Symbol),
	}
}

func (l *Linker) Link() error {
	l.putLinkerGlobals()

	err := l.putSegments()
	if err != nil {
		return err
	}

	err = l.resolveSymbols()
	if err != nil {
		return err
	}

	err = l.linkSymbols()
	if err != nil {
		return err
	}
	return nil
}

func (l *Linker) Out(path string) (int, error) {
	f, err := os.Create(path)
	if err != nil {
		return 0, fmt.Errorf("open file error. err: %s", err)
	}
	defer f.Close()

	err = binary.Write(f, binary.BigEndian, l.memory)
	if err != nil {
		return 0, fmt.Errorf("file write error. err: %s", err)
	}

	return len(l.memory), nil
}

func (l *Linker) resolveSymbols() error {
	for _, o := range l.objects {
		segment := o.Tracker.GetSegment()
		symbols := o.Tracker.GetSymbols()
		positions := o.Tracker.GetPositions()
		for _, p := range positions {
			sym := p.GetSymbol()
			s, ok := symbols[sym]
			if !ok {
				return fmt.Errorf("symbol is not found. forgot to tag as extern?. segment: %s, symbol: '%s'", segment, sym)
			}
			if p.IsMissing() {
				if s.IsShare(object.SYMBOL_SHARE_STATUS_GLOBAL) {
					return fmt.Errorf("symbol declared as 'global' but it is not found. segment: %s, symbol: '%s'", segment, sym)
				}
				// l.missing[segment] = s
				pushSymbol(segment, s, l.missing)
			}
		}
		for _, s := range symbols {
			if s.IsShare(object.SYMBOL_SHARE_STATUS_GLOBAL) {
				pushSymbol(segment, s, l.globals)
				// l.globals[segment] = s
				// fmt.Println("GLOBAL", segment, s.GetSymbol(), s.GetIndex(), s.GetType())
			}
			if s.IsShare(object.SYMBOL_SHARE_STATUS_EXTERN) {
				pushSymbol(segment, s, l.externs)
				// l.externs[segment] = s
				// fmt.Println("EXTERN", segment, s.GetSymbol(), s.GetIndex())
			}
		}
	}
	for segment, exs := range l.externs {
		for _, ex := range exs {
			externSymbol := ex.GetSymbol()
			_, _, ok := findGlobal(externSymbol, l.globals)
			if !ok {
				return fmt.Errorf("extern symbol not found. segment: %s, symbol: '%s'", segment, externSymbol)
			}
		}
	}
	return nil
}

func (l *Linker) putLinkerGlobals() {
	for _, m := range l.config.Memory {
		linkerGlobalSymbol := object.NewSymbol(fmt.Sprintf("__MEM_%s_START__", m.Name))
		linkerGlobalSymbol.SetIndex(m.Start)
		linkerGlobalSymbol.SetType(object.SYMBOL_TYPE_VAR)
		pushSymbol(m.Name, linkerGlobalSymbol, l.globals)
		// fmt.Println(linkerGlobalSymbol.GetSymbol(), linkerGlobalSymbol.GetIndex())

		linkerGlobalSymbol = object.NewSymbol(fmt.Sprintf("__MEM_%s_END__", m.Name))
		linkerGlobalSymbol.SetIndex(m.Start + m.Size)
		linkerGlobalSymbol.SetType(object.SYMBOL_TYPE_VAR)
		pushSymbol(m.Name, linkerGlobalSymbol, l.globals)
		// fmt.Println(linkerGlobalSymbol.GetSymbol(), linkerGlobalSymbol.GetIndex())
	}
}

func (l *Linker) putSegments() error {
	for _, s := range l.config.Segments {
		m, ok := l.config.GetMemoryConfig(s.Load)
		if !ok {
			return fmt.Errorf("memory config not found. segment: '%s', load: '%s'", s.Name, s.Load)
		}
		o, ok := findObjectFile(s.Name, l.objects)
		if !ok {
			return fmt.Errorf("segment not found in given object files. segment: '%s'", s.Name)
		}
		bin := o.Tracker.GetBin()
		length := uint16(len(bin))
		offset := l.lastMemoryOffsets[m.Name]
		position := m.Start + offset
		l.segmentOffsets[s.Name] = position
		// fmt.Printf("writing to %s to %s, from: %04x, to: %04x\n", s.Name, m.Name, position, position+length)
		copy(l.memory[position:], bin[:])
		l.lastMemoryOffsets[m.Name] += length
	}
	return nil
}

func (l *Linker) linkSymbols() error {
	for _, o := range l.objects {
		segment := o.Tracker.GetSegment()
		// fmt.Println("---", segment, "---")
		// symbols := o.Tracker.GetSymbols()
		positions := o.Tracker.GetPositions()
		for _, p := range positions {

			offset := uint16(0)
			index := uint16(0)
			sym := p.GetSymbol()
			segmentOffset := l.segmentOffsets[segment]
			offset = segmentOffset + p.GetOffset()
			if p.IsMissing() {
				globalSegment, globalSymbol, ok := findGlobal(sym, l.globals)
				if !ok {
					return fmt.Errorf("symbol not found. segment: %s, symbol: '%s'", segment, sym)
				}
				_type := globalSymbol.GetType()
				switch _type {
				case object.SYMBOL_TYPE_VAR:
					index = globalSymbol.GetIndex()
				case object.SYMBOL_TYPE_LABEL:
					globalSegmentOffset := l.segmentOffsets[globalSegment]
					index = globalSymbol.GetIndex() + globalSegmentOffset
				default:
					return fmt.Errorf("unknown symbol type. symbol: %s, type: %d", sym, _type)
				}
			} else {
				symbol, ok := o.Tracker.GetDefinedSymbol(sym)
				if !ok {
					return fmt.Errorf("fatal error object file claims have the symbol but not found in defined symbols. segment: %s, symbol: %s", segment, sym)
				}
				index = symbol.GetIndex() + segmentOffset
			}
			index += p.GetOptionalOffset()
			data := []byte{uint8(index), uint8(index >> 8)}
			size := uint16(1)
			if p.GetSize() == 16 {
				size = 2
			}
			copy(l.memory[offset:offset+size], data)
			// fmt.Printf("RESTORING LABEL >> object file segment: %s, missing_symbol_offset: %04x, global_symbol_segment: %s, global_symbol_type: %d, index(value): %04x\n", segment, offset, globalSegment, globalSymbol.GetType(), globalSymbol.GetIndex())
			// fmt.Printf("writing to segment %s, symbol: %s, from: %04x, to: %04x, data: %04x\n", segment, sym, offset, offset+size, data)
		}
	}
	return nil
}

func pushSymbol(segment string, s *object.Symbol, m map[string][]*object.Symbol) {
	_, ok := m[segment]
	if !ok {
		m[segment] = make([]*object.Symbol, 0, 4)
	}
	m[segment] = append(m[segment], s)
}

func findGlobal(symbol string, globals map[string][]*object.Symbol) (string, *object.Symbol, bool) {
	for segment, symbols := range globals {
		for _, sym := range symbols {
			if sym.GetSymbol() == symbol {
				return segment, sym, true
			}
		}
	}
	return "", nil, false
}

func findObjectFile(segment string, objects []*object.ELF) (*object.ELF, bool) {
	for _, o := range objects {
		if o.Tracker.GetSegment() == segment {
			return o, true
		}
	}
	return nil, false
}
