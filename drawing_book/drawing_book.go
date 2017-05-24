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
	var n, p int
	fmt.Fscanf(stdin, "%d %d", &n, &p)
	forward := p / 2
	var backward int
	if n%2 == 0 {
		backward = (n - p + 1) / 2
	} else {
		backward = (n - p) / 2
	}
	if backward < forward {
		fmt.Fprintln(stdout, backward)
	} else {
		fmt.Fprintln(stdout, forward)
	}
}
