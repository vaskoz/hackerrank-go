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
	var n, tallest, count int
	fmt.Fscanf(stdin, "%d", &n)
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscanf(stdin, "%d", &x)
		if x > tallest {
			tallest = x
			count = 1
		} else if x == tallest {
			count++
		}
	}
	fmt.Fprintln(stdout, count)
}
