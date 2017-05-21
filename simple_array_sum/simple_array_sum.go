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
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscanf(stdin, "%d", &x)
		sum += x
	}
	fmt.Fprintln(stdout, sum)
}
