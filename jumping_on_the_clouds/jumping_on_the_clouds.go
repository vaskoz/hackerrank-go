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
	var clouds, result, x int
	fmt.Fscanf(stdin, "%d", &clouds)
	steps := -1
	for i := 0; i < clouds; i++ {
		fmt.Fscanf(stdin, "%d", &x)
		if x == 0 {
			steps++
		} else {
			result += (steps / 2) + (steps % 2) + 1
			steps = -1
		}
	}
	result += (steps / 2) + (steps % 2)
	fmt.Fprintln(stdout, result)
}
