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
	var p, q, answers int
	fmt.Fscanf(stdin, "%d", &p)
	fmt.Fscanf(stdin, "%d", &q)
	for i := p; i <= q; i++ {
		x := i * i
		d := digits(i)
		l, r := split(x, d)
		if l+r == i {
			answers++
			fmt.Fprintf(stdout, "%d ", i)
		}
	}
	if answers == 0 {
		fmt.Fprint(stdout, "INVALID RANGE")
	}
	fmt.Fprintln(stdout)
}

func digits(x int) int {
	result := 1
	for x > 0 {
		x /= 10
		result *= 10
	}
	return result
}

func split(x, d int) (l, r int) {
	l, r = 0, 0
	r = x % d
	x /= d
	l = x
	return l, r
}
