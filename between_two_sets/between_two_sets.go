package main

import (
	"fmt"
	"io"
	"os"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * (b / gcd(a, b))
}

func main() {
	var n, m int
	fmt.Fscanf(stdin, "%d %d", &n, &m)
	var f int
	fmt.Fscanf(stdin, "%d", &f)
	for i := 1; i < n; i++ {
		var x int
		fmt.Fscanf(stdin, "%d", &x)
		f = lcm(f, x)
	}

	var l int
	fmt.Fscanf(stdin, "%d", &l)
	for i := 1; i < m; i++ {
		var x int
		fmt.Fscanf(stdin, "%d", &x)
		l = gcd(l, x)
	}

	count := 0
	for i, j := f, 2; i <= l; i, j = f*j, j+1 {
		if l%i == 0 {
			count++
		}
	}
	fmt.Fprintln(stdout, count)
}
