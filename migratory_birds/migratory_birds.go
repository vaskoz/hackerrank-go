package main

import (
	"fmt"
	"io"
	"os"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout

func main() {
	var n int
	fmt.Fscanf(stdin, "%d", &n)
	var types [5]int
	for i := 0; i < n; i++ {
		var t int
		fmt.Fscanf(stdin, "%d", &t)
		types[t-1]++
	}
	max := -1
	maxi := -1
	for i := 0; i < len(types); i++ {
		if types[i] > max {
			max = types[i]
			maxi = i
		}
	}
	fmt.Fprintln(stdout, maxi+1)
}
