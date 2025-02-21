package test

import (
	"testing"
)

type TestCase struct {
	Name     string
	Code     string
	Expected []byte
	Debug    bool
}

func RunCase(t *testing.T, tc *TestCase) {
	t.Run(tc.Name, func(tt *testing.T) {
		out, err := Assemble(tc.Code)
		if err != nil {
			tt.Fatalf("assemble failed. err: %s", err.Error())
		}
		OutTest(tt, tc.Expected, out)
	})
}
