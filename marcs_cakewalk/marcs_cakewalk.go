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
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	result := 0
	twos := 1
	for _, c := range a {
		result += c * twos
		twos *= 2
	}
	fmt.Fprintln(stdout, result)
}
