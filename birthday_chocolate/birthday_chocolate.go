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
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(stdin, "%d", &s[i])
	}
	var d, m, result int
	fmt.Fscanf(stdin, "%d %d", &d, &m)
	for i := 0; i <= n-m; i++ {
		sum := 0
		for j := i; j < i+m; j++ {
			sum += s[j]
		}
		if sum == d {
			result++
		}
	}
	fmt.Fprintln(stdout, result)
}
