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
	var n, k, diff, height int
	fmt.Fscanf(stdin, "%d %d", &n, &k)
	for i := 0; i < n; i++ {
		fmt.Fscanf(stdin, "%d", &height)
		if height > k && height-k > diff {
			diff = height - k
		}
	}
	fmt.Fprintln(stdout, diff)
}
