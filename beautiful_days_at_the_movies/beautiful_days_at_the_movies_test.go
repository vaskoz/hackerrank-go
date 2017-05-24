package main

import (
	"bytes"
	"strings"
	"testing"
)

var testcases = []struct {
	in, out string
}{
	{"20 23 6", "2\n"},
}

func TestBeautifulDaysAtTheMovies(t *testing.T) {
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

func BenchmarkBeautifulDaysAtTheMovies(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			stdin = strings.NewReader(c.in)
			buff := new(bytes.Buffer)
			stdout = buff
			main()
		}
	}
}

var reverseTestcases = []struct {
	in, out int
}{
	{1234, 4321},
	{42341234, 43214324},
	{10, 1},
	{9, 9},
}

func TestReverseUsingInt(t *testing.T) {
	for _, c := range reverseTestcases {
		if out := ReverseUsingInt(c.in); out != c.out {
			t.Errorf("Expected %v, but got %v", c.out, out)
		}
	}
}

func TestReverseUsingString(t *testing.T) {
	for _, c := range reverseTestcases {
		if out := ReverseUsingString(c.in); out != c.out {
			t.Errorf("Expected %v, but got %v", c.out, out)
		}
	}
}

func BenchmarkReverseUsingInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range reverseTestcases {
			ReverseUsingInt(c.in)
		}
	}
}

func BenchmarkReverseUsingString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range reverseTestcases {
			ReverseUsingString(c.in)
		}
	}
}
