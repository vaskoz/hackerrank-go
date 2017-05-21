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
	var s, t, a, b, m, n, appleHits, orangeHits, pos int
	fmt.Fscanf(stdin, "%d %d", &s, &t)
	fmt.Fscanf(stdin, "%d %d", &a, &b)
	fmt.Fscanf(stdin, "%d %d", &m, &n)
	for i := 0; i < m; i++ {
		fmt.Fscanf(stdin, "%d", &pos)
		if a+pos >= s && a+pos <= t {
			appleHits++
		}
	}
	for i := 0; i < n; i++ {
		fmt.Fscanf(stdin, "%d", &pos)
		if b+pos >= s && b+pos <= t {
			orangeHits++
		}
	}
	fmt.Fprintf(stdout, "%d\n%d\n", appleHits, orangeHits)
}
