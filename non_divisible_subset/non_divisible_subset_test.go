package main

import (
	"bytes"
	"strings"
	"testing"
)

var testcases = []struct {
	in, out string
}{
	{"4 3 1 7 2 4", "3\n"},
	{"10 5 770528134 663501748 384261537 800309024 103668401 538539662 385488901 101262949 557792122 46058493", "6\n"},
	{"6 9 422346306 940894801 696810740 862741861 85835055 313720373", "5\n"},
	{"3 10 10 20 30", "1\n"},
	{"3 10 10 20 35", "2\n"},
}

func TestNonDivisibleSubset(t *testing.T) {
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

func BenchmarkNonDivisibleSubset(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			stdin = strings.NewReader(c.in)
			buff := new(bytes.Buffer)
			stdout = buff
			main()
		}
	}
}
