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
	var x1, v1, x2, v2 int
	fmt.Fscanf(stdin, "%d %d %d %d", &x1, &v1, &x2, &v2)
	deltaX := x2 - x1
	deltaV := v1 - v2
	if deltaV != 0 && deltaX%deltaV == 0 && deltaX/deltaV > 0 {
		fmt.Fprintln(stdout, "YES")
	} else {
		fmt.Fprintln(stdout, "NO")
	}
}
