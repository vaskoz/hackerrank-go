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
	var n int
	fmt.Fscanf(stdin, "%d", &n)
	sticks := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(stdin, "%d", &sticks[i])
	}
	sort.Ints(sticks) // O(N*log(N))
	last := 0
	for i, s := range sticks { // O(N)
		if s != last {
			last = s
			fmt.Fprintln(stdout, len(sticks)-i)
		}
	}
}
