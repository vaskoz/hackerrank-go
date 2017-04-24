package main

import (
	"fmt"
	"io"
	"os"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout

func main() {
	var n int
	fmt.Fscanf(stdin, "%d", &n)
	var s string
	fmt.Fscanf(stdin, "%s", &s)
	inValley := false
	result := 0
	pos := 0
	for _, r := range s {
		if r == 'D' {
			pos--
		} else if r == 'U' {
			pos++
		}
		if inValley && pos == 0 {
			inValley = false
			result++
		} else if !inValley && pos < 0 {
			inValley = true
		}
	}
	fmt.Fprintln(stdout, result)
}
