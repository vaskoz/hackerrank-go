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
	sum := 0
	sumpos, negpos := -1, -1
	for i := 0; i < n*n; i++ {
		var x int
		fmt.Fscanf(stdin, "%d", &x)
		if i%n == 0 {
			sumpos++
			negpos++
		}
		if i%n == sumpos {
			sum += x
		}
		if i%n == n-negpos-1 {
			sum -= x
		}
	}
	if sum < 0 {
		sum = -sum
	}
	fmt.Fprintln(stdout, sum)
}
