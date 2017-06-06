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
	var n, x, result int
	fmt.Fscanf(stdin, "%d", &n)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Fscanf(stdin, "%d", &x)
		m[x]++
	}
	maxFreq, maxKey := 0, 0
	for k, v := range m {
		if v > maxFreq {
			maxFreq, maxKey = v, k
		}
	}
	delete(m, maxKey)
	for _, v := range m {
		result += v
	}
	fmt.Fprintln(stdout, result)
}
