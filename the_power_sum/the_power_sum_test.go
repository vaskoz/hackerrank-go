package main

import (
	"bytes"
	"strings"
	"testing"
)

var testcases = []struct {
	in, out string
}{
	{"5 2", "1\n"},
	{"4 2", "1\n"},
	{"6 2", "0\n"},
	{"10 2", "1\n"},
	{"100 2", "3\n"},
	{"100 3", "1\n"},
	{"1000 10", "0\n"},
}

var powerTestcases = []struct {
	a, b, out int
}{
	{1, 3, 1},
	{2, 2, 4},
	{10, 2, 100},
}

func TestThePowerSum(t *testing.T) {
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

func BenchmarkThePowerSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			stdin = strings.NewReader(c.in)
			buff := new(bytes.Buffer)
			stdout = buff
			main()
		}
	}
}

func TestPower(t *testing.T) {
	for _, c := range powerTestcases {
		if out := Power(c.a, c.b); out != c.out {
			t.Errorf("For %d to the %d should be %d, but got %d", c.a, c.b, c.out, out)
		}
	}
}

func BenchmarkPower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range powerTestcases {
			Power(c.a, c.b)
		}
	}
}
