package main

import (
	"fmt"
	"io"
	"os"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr

func main() {
	var testCases, n, m, s int
	fmt.Fscanf(stdin, "%d", &testCases)
	for test := 0; test < testCases; test++ {
		fmt.Fscanf(stdin, "%d %d %d", &n, &m, &s)
		if pos := (s - 1 + m) % n; pos == 0 {
			fmt.Fprintln(stdout, n)
		} else {
			fmt.Fprintln(stdout, pos)
		}
	}
}
