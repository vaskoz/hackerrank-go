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
	var n, result int
	fmt.Fscanf(stdin, "%d", &n)
	dayTotal := 5
	for i := 0; i < n; i++ {
		dayTotal = (dayTotal / 2)
		result += dayTotal
		dayTotal *= 3
	}
	fmt.Fprintln(stdout, result)
}
