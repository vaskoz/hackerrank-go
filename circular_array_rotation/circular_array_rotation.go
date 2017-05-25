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
	var n, k, q int
	fmt.Fscanf(stdin, "%d %d %d", &n, &k, &q)
	input := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(stdin, "%d", &input[i])
	}
	var query int
	rotation := k % n
	for i := 0; i < q; i++ {
		fmt.Fscanf(stdin, "%d", &query)
		fmt.Fprintln(stdout, input[(n-rotation+query)%n])
	}
}
