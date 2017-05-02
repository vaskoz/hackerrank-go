package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout

func main() {
	var n int
	fmt.Fscanf(stdin, "%d", &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(stdin, "%d", &a[i])
	}
	sort.Ints(a)
	minimum := a[1] - a[0]
	for i := 1; i < n; i++ {
		if diff := a[i] - a[i-1]; diff < minimum {
			minimum = diff
		}
	}
	fmt.Fprintln(stdout, minimum)
}
