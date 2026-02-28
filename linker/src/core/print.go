package core

import (
	"asme8/common/object"
	"fmt"
	"sort"

	"github.com/ecoshub/stable"
)

type ResolvedSymbol struct {
	segment        string
	symbol         string
	_type          uint8
	index          uint16
	optionalOffset uint16
}

func (l *Linker) PrintSymbols(printCode bool) {

	fmt.Println()

	t := stable.New("GLOBAL LINKER SYMBOLS")
	t.AddFieldWithOptions("ADDR", &stable.Options{
		Format:          "$%04x",
		Alignment:       stable.AlignmentLeft,
		HeaderAlignment: stable.AlignmentLeft,
	})
	t.AddFieldWithOptions("SYMBOL", &stable.Options{
		Format:          "%s",
		Alignment:       stable.AlignmentLeft,
		HeaderAlignment: stable.AlignmentLeft,
	})
	for _, s := range l.linkerSymbols {
		t.Row(s.GetIndex(), s.GetSymbol())
	}
	fmt.Println(t)

	t = stable.New("RESOLVED SYMBOLS")
	t.AddFieldWithOptions("SEGMENT", &stable.Options{
		Format:          "%s",
		Alignment:       stable.AlignmentLeft,
		HeaderAlignment: stable.AlignmentLeft,
	})
	t.AddFieldWithOptions("TYPE", &stable.Options{
		Format:          "%s",
		Alignment:       stable.AlignmentCenter,
		HeaderAlignment: stable.AlignmentCenter,
	})
	t.AddFieldWithOptions("ADDR", &stable.Options{
		Format:          "$%04x",
		Alignment:       stable.AlignmentLeft,
		HeaderAlignment: stable.AlignmentLeft,
	})
	t.AddFieldWithOptions("SYMBOL", &stable.Options{
		Format:          "%s",
		Alignment:       stable.AlignmentLeft,
		HeaderAlignment: stable.AlignmentLeft,
	})

	for _, seg := range l.config.SegmentConfig.Configs {
		rss := l.resolvedSymbols[seg.Name]
		unique := map[string]*ResolvedSymbol{}
		for _, rs := range rss {
			if rs.optionalOffset != 0 {
				continue
			}
			unique[rs.symbol] = rs
		}
		sorted := SortSymbolMap(unique)
		for _, us := range sorted {
			us := unique[us]
			typ := "NON"
			if us._type == object.SYMBOL_TYPE_VAR {
				typ = "VAR"
			}
			if us._type == object.SYMBOL_TYPE_LABEL {
				typ = "LAB"
			}
			t.Row(us.segment, typ, us.index, us.symbol)
		}
	}
	fmt.Println(t)

	if printCode {
		t = stable.New("")
		t.AddFieldWithOptions("ADDR", &stable.Options{
			Alignment:       stable.AlignmentLeft,
			HeaderAlignment: stable.AlignmentLeft,
		})
		t.AddFieldWithOptions("CODE", &stable.Options{
			Alignment:       stable.AlignmentLeft,
			HeaderAlignment: stable.AlignmentLeft,
			CharLimit:       100,
		})
		for _, line := range l.code {
			t.Row(line[:6], line[6:])
		}

		fmt.Println(t)
	}

}

func SortSymbolMap(m map[string]*ResolvedSymbol) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]].index < m[keys[j]].index
	})

	return keys
}
