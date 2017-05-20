package main

import (
	"bytes"
	"strings"
	"testing"
)

var testcases = []struct {
	in  string
	out string
}{
	{"4 1 2 3 4", "3 0\n"},
	{"10 3 4 21 36 10 28 35 5 24 42", "4 0\n"},
	{"9 10 5 20 20 4 5 2 25 1", "2 4\n"},
	{"4 4 3 2 1", "0 3\n"},
}

func TestBestWorst(t *testing.T) {
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

func BenchmarkBestWorst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			stdin = strings.NewReader(c.in)
			buff := new(bytes.Buffer)
			stdout = buff
			main()
		}
	}
}
