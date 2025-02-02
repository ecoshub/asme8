package core

import (
	"fmt"
	"sort"

	"github.com/ecoshub/stable"
)

type ResolvedSymbol struct {
	segment string
	symbol  string
	index   uint16
}

func (l *Linker) PrintSymbols(printCode bool) {

	fmt.Println()

	t := stable.New("GLOBAL LINKER SYMBOLS")
	t.AddFieldWithOptions("INDEX", &stable.Options{
		Format:          "%04x",
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
	t.AddFieldWithOptions("INDEX", &stable.Options{
		Format:          "%04x",
		Alignment:       stable.AlignmentLeft,
		HeaderAlignment: stable.AlignmentLeft,
	})
	t.AddFieldWithOptions("SYMBOL", &stable.Options{
		Format:          "%s",
		Alignment:       stable.AlignmentLeft,
		HeaderAlignment: stable.AlignmentLeft,
	})

	for _, seg := range l.segmentConfig.Configs {
		rss := l.resolvedSymbols[seg.Name]
		unique := map[string]*ResolvedSymbol{}
		for _, rs := range rss {
			unique[rs.symbol] = rs
		}
		sorted := SortSymbolMap(unique)
		for _, us := range sorted {
			us := unique[us]
			t.Row(us.segment, us.index, us.symbol)
		}
	}
	fmt.Println(t)

	if printCode {
		t = stable.New("")
		t.AddFieldWithOptions("OFFSET", &stable.Options{
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
