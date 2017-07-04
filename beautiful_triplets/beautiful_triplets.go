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
	var n, d, result int
	fmt.Fscanf(stdin, "%d %d", &n, &d)
	a := make([]int, n)
	lookup := make(map[int]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(stdin, "%d", &a[i])
		lookup[a[i]] = i
	}
	for i := 0; i < n-2; i++ {
		aj := d + a[i]
		if j, found := lookup[aj]; found && i < j {
			ak := d + a[j]
			if k, found := lookup[ak]; found && j < k {
				result++
			}
		}
	}
	fmt.Fprintln(stdout, result)
}
