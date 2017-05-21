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
	var n int
	fmt.Fscanf(stdin, "%d", &n)
	var positive, negative, zeros float64
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscanf(stdin, "%d", &x)
		if x > 0 {
			positive += 1
		} else if x == 0 {
			zeros += 1
		} else {
			negative += 1
		}
	}
	nf := float64(n)
	fmt.Fprintf(stdout, "%.6f\n%.6f\n%.6f\n", positive/nf, negative/nf, zeros/nf)
}
