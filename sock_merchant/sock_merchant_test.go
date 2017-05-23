package main

import (
	"bytes"
	"strings"
	"testing"
)

var testcases = []struct {
	in, out string
}{
	{"9 10 20 20 10 10 30 50 10 20", "3\n"},
}

func TestSockMerchant(t *testing.T) {
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

func BenchmarkSockMerchant(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			stdin = strings.NewReader(c.in)
			buff := new(bytes.Buffer)
			stdout = buff
			main()
		}
	}
}
