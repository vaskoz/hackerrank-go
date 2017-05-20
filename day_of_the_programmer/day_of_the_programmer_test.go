package main

import (
	"bytes"
	"strings"
	"testing"
)

var testcases = []struct {
	in, out string
}{
	{"2017", "13.09.2017\n"},
	{"2016", "12.09.2016\n"},
	{"2025", "13.09.2025\n"},
	{"1918", "26.09.1918\n"},
	{"1912", "12.09.1912\n"},
	{"1911", "13.09.1911\n"},
}

func TestDayOfTheProgrammer(t *testing.T) {
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

func BenchmarkDayOfTheProgrammer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			stdin = strings.NewReader(c.in)
			buff := new(bytes.Buffer)
			stdout = buff
			main()
		}
	}
}
