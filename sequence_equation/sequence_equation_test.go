package main

import (
	"bytes"
	"strings"
	"testing"
)

var testcases = []struct {
	in, out string
}{
	{"3 2 3 1", "2\n3\n1\n"},
}

func TestSequenceEquation(t *testing.T) {
	t.Parallel()
	for _, c := range testcases {
		stdin = strings.NewReader(c.in)
		buff := new(bytes.Buffer)
		stdout = buff
		main()
		out := buff.String()
		if out != c.out {
			t.Errorf("%v does not equal %v", out, c.out)
		}
	}
}

func BenchmarkSequenceEquation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			stdin = strings.NewReader(c.in)
			buff := new(bytes.Buffer)
			stdout = buff
			main()
		}
	}
}
