package utils

import "fmt"

type FlagArray []string

// String is an implementation of the flag.Value interface
func (i *FlagArray) String() string {
	return fmt.Sprintf("%v", *i)
}

// Set is an implementation of the flag.Value interface
func (i *FlagArray) Set(value string) error {
	*i = append(*i, value)
	return nil
}
