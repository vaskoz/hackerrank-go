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
	var tests, n, t int64
	fmt.Fscanf(stdin, "%d", &tests)
	for t = 0; t < tests; t++ {
		result := 0
		fmt.Fscanf(stdin, "%d", &n)
		temp := n
		for temp != 0 {
			d := temp % 10
			if d != 0 && n%d == 0 {
				result++
			}
			temp /= 10
		}
		fmt.Fprintln(stdout, result)
	}
}
