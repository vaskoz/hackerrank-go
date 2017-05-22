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
	var n, k, expected, charged int
	fmt.Fscanf(stdin, "%d %d", &n, &k)
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscanf(stdin, "%d", &x)
		if i != k {
			expected += x
		}
	}
	fmt.Fscanf(stdin, "%d", &charged)
	if charged*2 == expected {
		fmt.Fprintln(stdout, "Bon Appetit")
	} else {
		fmt.Fprintln(stdout, charged-expected/2)
	}
}
