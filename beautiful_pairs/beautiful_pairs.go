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
	mapA := make(map[int]int, n)
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscanf(stdin, "%d", &x)
		mapA[x]++
	}
	result := 0
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscanf(stdin, "%d", &x)
		if count, found := mapA[x]; found {
			result++
			if count == 1 {
				delete(mapA, x)
			} else {
				mapA[x]--
			}
		}
	}
	if result != n {
		fmt.Fprintln(stdout, result+1)
	} else {
		fmt.Fprintln(stdout, result-1)
	}
}
