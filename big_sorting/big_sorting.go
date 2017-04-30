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
	a := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(stdin, "%s", &a[i])
	}
	sort.Slice(a, func(i, j int) bool {
		if al, bl := len(a[i]), len(a[j]); al < bl {
			return true
		} else if al == bl {
			return a[i] < a[j]
		} else {
			return false
		}
	})
	for i := 0; i < n; i++ {
		fmt.Fprintln(stdout, a[i])
	}
}
