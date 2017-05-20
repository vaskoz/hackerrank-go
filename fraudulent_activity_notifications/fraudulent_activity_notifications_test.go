package main

import (
	"bytes"
	"strings"
	"testing"
)

var testcases = []struct {
	in, out string
}{
	{"9 5 2 3 4 2 3 6 8 4 5", "2\n"},
	{"5 4 1 2 3 4 4", "0\n"},
	{"5 2 5 4 3 2 1", "0\n"},
	{"10 4 10 9 8 7 1 5 4 3 2 11", "1\n"},
}

func TestFraudulentActivityNotifications(t *testing.T) {
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

func BenchmarkFraudulentActivityNotifications(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			stdin = strings.NewReader(c.in)
			buff := new(bytes.Buffer)
			stdout = buff
			main()
		}
	}
}
