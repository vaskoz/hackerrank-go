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
		var b, w int
		fmt.Fscanf(stdin, "%d %d", &b, &w)
		var x, y, z int
		fmt.Fscanf(stdin, "%d %d %d", &x, &y, &z)
		result := 0
		if x+z < y {
			result = b*x + w*(x+z)
		} else if y+z < x {
			result = b*(y+z) + w*y
		} else {
			result = b*x + w*y
		}
		fmt.Fprintln(stdout, result)
	}
}
