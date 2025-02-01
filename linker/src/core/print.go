package core

import (
	"fmt"
	"sort"
)

type ResolvedSymbol struct {
	segment string
	symbol  string
	index   uint16
}

func (l *Linker) PrintSymbols() {
	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Println("RESOLVED SYMBOLS:")
	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Println("SEGMENT                       INDEX       SYMBOL")
	fmt.Println("--------------------------------------------------------------------------------------------")
	for _, seg := range l.config.Segments {
		rss := l.resolvedSymbols[seg.Name]
		unique := map[string]*ResolvedSymbol{}
		for _, rs := range rss {
			unique[rs.symbol] = rs
		}
		sorted := SortSymbolMap(unique)
		for _, us := range sorted {
			us := unique[us]
			fmt.Printf("%-20s      %04x        <%s>\n", us.segment, us.index, us.symbol)
		}
	}
	fmt.Println("--------------------------------------------------------------------------------------------")
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
