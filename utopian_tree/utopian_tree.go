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
	var tests int
	fmt.Fscanf(stdin, "%d", &tests)
	for t := 0; t < tests; t++ {
		height := 1
		var cycles int
		fmt.Fscanf(stdin, "%d", &cycles)
		for i := 0; i < cycles; i++ {
			if i%2 == 0 {
				height *= 2
			} else {
				height += 1
			}
		}
		fmt.Fprintln(stdout, height)
	}
}
