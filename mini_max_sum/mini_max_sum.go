package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr

func main() {
	a := make([]int, 5)
	sum := 0
	for i := range a {
		fmt.Fscanf(stdin, "%d", &a[i])
		sum += a[i]
	}
	sort.Ints(a)
	fmt.Fprintln(stdout, sum-a[len(a)-1], sum-a[0])
}
