package main

import (
	"bytes"
	"strings"
	"testing"
)

var testcases = []struct {
	in, out string
}{
	{"4 5 10101 11100 11010 00101", "5\n2\n"},
}

func TestAcmIcpcTeam(t *testing.T) {
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

func BenchmarkAcmIcpcTeam(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			stdin = strings.NewReader(c.in)
			buff := new(bytes.Buffer)
			stdout = buff
			main()
		}
	}
}
