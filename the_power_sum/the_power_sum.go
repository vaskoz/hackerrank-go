package main

import (
	"fmt"
	"io"
	"os"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout

func main() {
	var x int
	fmt.Fscanf(stdin, "%d", &x)
	var n int
	fmt.Fscanf(stdin, "%d", &n)
	fmt.Fprintln(stdout, PowerSum(0, 1, n, x))
}

func PowerSum(sum, startWith, n, x int) int {
	if sum == x {
		return 1
	} else {
		count := 0
		for start := startWith; start < x; start++ {
			if newSum := sum + Power(start, n); newSum > x {
				break
			} else {
				count += PowerSum(newSum, start+1, n, x)
			}
		}
		return count
	}
}

func Power(a, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}
