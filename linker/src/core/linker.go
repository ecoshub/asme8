package core

import (
	"asme8/common/config"
	"asme8/common/object"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
)

type Linker struct {
	config                  *config.Config
	objects                 []*object.ELF
	memory                  []byte
	code                    []string
	memoryLastSegmentOffset map[string]uint16
	segmentOffsets          map[string]uint16
	globals                 map[string][]*object.Symbol
	linkerSymbols           []*object.Symbol
	externs                 map[string][]*object.Symbol
	missing                 map[string][]*object.Symbol
	resolvedSymbols         map[string][]*ResolvedSymbol
}

func NewLinker(conf *config.Config, objects ...*object.ELF) *Linker {
	return &Linker{
		config:                  conf,
		objects:                 objects,
		memory:                  make([]byte, 0xffff),
		memoryLastSegmentOffset: make(map[string]uint16),
		segmentOffsets:          make(map[string]uint16),
		linkerSymbols:           make([]*object.Symbol, 0, 4),
		globals:                 make(map[string][]*object.Symbol),
		externs:                 make(map[string][]*object.Symbol),
		missing:                 make(map[string][]*object.Symbol),
		resolvedSymbols:         make(map[string][]*ResolvedSymbol),
	}
}

func (l *Linker) Link() error {

	err := ResolveMemoryLayout(l.config.MemoryConfig.Configs)
	if err != nil {
		return err
	}

	l.putLinkerGlobals()

	err = l.putSegments()
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

	err = l.ResolveCode()
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

func (l *Linker) OutCode(path string) (int, error) {
	f, err := os.Create(path)
	if err != nil {
		return 0, fmt.Errorf("open file error. err: %s", err)
	}
	defer f.Close()

	err = os.WriteFile(path, []byte(strings.Join(l.code, "\n")), 0644)
	if err != nil {
		return 0, fmt.Errorf("file write error. err: %s", err)
	}

	return len(l.memory), nil
}

func (l *Linker) putLinkerGlobals() {
	for _, m := range l.config.MemoryConfig.Configs {
		linkerGlobalSymbol := object.NewSymbol(fmt.Sprintf("__%s_START__", m.Name))
		linkerGlobalSymbol.SetIndex(m.Start.Value)
		linkerGlobalSymbol.SetType(object.SYMBOL_TYPE_VAR)
		pushSymbol(m.Name, linkerGlobalSymbol, l.globals)
		l.linkerSymbols = append(l.linkerSymbols, linkerGlobalSymbol)
		linkerGlobalSymbol = object.NewSymbol(fmt.Sprintf("__%s_END__", m.Name))
		linkerGlobalSymbol.SetIndex(m.Start.Value + m.Size)
		linkerGlobalSymbol.SetType(object.SYMBOL_TYPE_VAR)
		pushSymbol(m.Name, linkerGlobalSymbol, l.globals)
		l.linkerSymbols = append(l.linkerSymbols, linkerGlobalSymbol)
	}
}

func ResolveMemoryLayout(ml []*config.Memory) error {
	offset := uint16(0)
	for _, m := range ml {
		if m.Start == nil {
			m.Start = &config.NullableValue{Value: offset}
			offset += m.Size
		}
	}
	return nil
}

func (l *Linker) putSegments() error {
	for _, s := range l.config.SegmentConfig.Configs {
		if s.Start == nil {
			continue
		}
		m, _ := l.config.MemoryConfig.GetMemoryConfig(s.Load)
		o, _ := findObjectFile(s.Name, l.objects)
		l.putImplicitSegments(o, m, s)
	}
	for _, s := range l.config.SegmentConfig.Configs {
		if s.Start != nil {
			continue
		}
		m, ok := l.config.MemoryConfig.GetMemoryConfig(s.Load)
		if !ok {
			return fmt.Errorf("memory config not found. segment: '%s', load: '%s'", s.Name, s.Load)
		}
		o, ok := findObjectFile(s.Name, l.objects)
		if !ok {
			return fmt.Errorf("segment not found in given object files. segment: '%s'", s.Name)
		}
		l.putExplicitSegment(o, m, s)
	}
	return nil
}

func (l *Linker) putImplicitSegments(o *object.ELF, m *config.Memory, s *config.Segment) error {
	bin := o.Tracker.GetBin()
	length := uint16(len(bin))
	position := m.Start.Value + s.Start.Value
	if position+length > m.Start.Value+m.Size {
		return fmt.Errorf("segment overflow. memory: %s, segment: %s", m.Name, s.Load)
	}
	l.segmentOffsets[s.Name] = position
	copy(l.memory[position:position+length], bin[:length])
	return nil
}

func (l *Linker) putExplicitSegment(o *object.ELF, m *config.Memory, s *config.Segment) error {
	bin := o.Tracker.GetBin()
	length := uint16(len(bin))
	offset := l.memoryLastSegmentOffset[m.Name]
	position := m.Start.Value + offset
	l.segmentOffsets[s.Name] = position
	if position+length > m.Start.Value+m.Size {
		return fmt.Errorf("segment overflow. memory: %s, segment: %s", m.Name, s.Load)
	}
	copy(l.memory[position:position+length], bin[:length])
	l.memoryLastSegmentOffset[m.Name] += length
	return nil
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
				pushSymbol(segment, s, l.missing)
			}
		}
		for _, s := range symbols {
			if s.IsShare(object.SYMBOL_SHARE_STATUS_GLOBAL) {
				existsSymbol, existsSegment, exists := checkSameGlobal(s.GetSymbol(), l.globals)
				if exists {
					return fmt.Errorf("symbol is already defined in another segment. symbol: %s, segment: %s, active segment: %s", existsSymbol.GetSymbol(), existsSegment, segment)
				}
				pushSymbol(segment, s, l.globals)
			}
			if s.IsShare(object.SYMBOL_SHARE_STATUS_EXTERN) {
				pushSymbol(segment, s, l.externs)
			}
		}
	}
	err := l.resolveReference()
	if err != nil {
		return err
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

func (l *Linker) resolveReference() error {
	for _, o := range l.objects {
		segment := o.Tracker.GetSegment()
		_, exists := l.config.SegmentConfig.GetSegmentConfig(segment)
		if !exists {
			continue
		}
		symbols := o.Tracker.GetSymbols()
		for _, s := range symbols {
			ref, offset := s.GetReference()
			if ref == "" {
				continue
			}
			segment, sym, exists := findGlobal(ref, l.globals)
			if exists {
				s.SetIndex(uint16(int32(sym.GetIndex()) + offset))
				s.SetReference("", 0)
				s.SetType(sym.GetType())
				pushSymbol(segment, s, l.globals)
				continue
			}
			return fmt.Errorf("reference symbol not found. segment: %s, symbol: %s, reference: %s", segment, s.GetSymbol(), ref)
		}
	}
	return nil
}

func (l *Linker) linkSymbols() error {
	for _, o := range l.objects {
		segment := o.Tracker.GetSegment()
		sc, exists := l.config.SegmentConfig.GetSegmentConfig(segment)
		if !exists {
			continue
		}

		// symbols := o.Tracker.GetSymbols()
		positions := o.Tracker.GetPositions()
		for _, p := range positions {
			offset := uint16(0)
			index := uint16(0)
			sym := p.GetSymbol()
			segmentOffset := l.segmentOffsets[segment]
			offset = segmentOffset + sc.Start.Get() + p.GetOffset()
			globalSegmentOffset := uint16(0)
			_type := uint8(0)
			if p.IsMissing() {
				globalSegment, globalSymbol, ok := findGlobal(sym, l.globals)
				if !ok {
					return fmt.Errorf("symbol not found. segment: %s, symbol: '%s'", segment, sym)
				}
				_type = globalSymbol.GetType()
				switch _type {
				case object.SYMBOL_TYPE_VAR:
					index = globalSymbol.GetIndex()
				case object.SYMBOL_TYPE_LABEL:
					globalSegmentOffset = l.segmentOffsets[globalSegment]
					index = globalSymbol.GetIndex() + globalSegmentOffset
				default:
					return fmt.Errorf("unknown symbol type for missing symbol. segment: %s, symbol: %s, type: %d", segment, sym, _type)
				}
			} else {
				symbol, ok := o.Tracker.GetDefinedSymbol(sym)
				if !ok {
					return fmt.Errorf("fatal error object file claims have the symbol but not found in defined symbols. segment: %s, symbol: %s", segment, sym)
				}
				_type = symbol.GetType()
				switch _type {
				case object.SYMBOL_TYPE_VAR:
					index = symbol.GetIndex()
				case object.SYMBOL_TYPE_LABEL:
					globalSegmentOffset = l.segmentOffsets[segment]
					index = symbol.GetIndex() + globalSegmentOffset
				default:
					return fmt.Errorf("unknown symbol type. segment: %s, symbol: %s, type: %d", segment, sym, _type)
				}
			}
			index += p.GetOptionalOffset()
			data := []byte{uint8(index), uint8(index >> 8)}
			size := uint16(1)
			if p.GetSize() == 16 {
				size = 2
			}
			copy(l.memory[offset:offset+size], data)
			if p.IsMissing() {
				rs := &ResolvedSymbol{segment: segment, _type: _type, symbol: sym, index: index, optionalOffset: p.GetOptionalOffset()}
				pushResolvedSymbol(segment, rs, l.resolvedSymbols)
			}
		}
	}
	return nil
}

func pushResolvedSymbol(segment string, s *ResolvedSymbol, m map[string][]*ResolvedSymbol) {
	_, ok := m[segment]
	if !ok {
		m[segment] = make([]*ResolvedSymbol, 0, 4)
	}
	m[segment] = append(m[segment], s)
}

func pushSymbol(segment string, s *object.Symbol, m map[string][]*object.Symbol) {
	_, ok := m[segment]
	if !ok {
		m[segment] = make([]*object.Symbol, 0, 4)
	}
	m[segment] = append(m[segment], s)
}

func checkSameGlobal(symbol string, globals map[string][]*object.Symbol) (*object.Symbol, string, bool) {
	for segment, arr := range globals {
		for _, s := range arr {
			if s.GetSymbol() == symbol {
				return s, segment, true
			}
		}
	}
	return nil, "", false
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
