package main

import (
	"bytes"
	"strings"
	"testing"
)

var testcases = []struct {
	in, out string
}{
	{"5 ab bb hefg dhck dkhc", "ba\nno answer\nhegf\ndhkc\nhcdk\n"},
	{"1 a", "no answer\n"},
}

func TestBiggerIsGreater(t *testing.T) {
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

func BenchmarkBiggerIsGreater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			stdin = strings.NewReader(c.in)
			buff := new(bytes.Buffer)
			stdout = buff
			main()
		}
	}
}
