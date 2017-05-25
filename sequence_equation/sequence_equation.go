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
	var n, x int
	fmt.Fscanf(stdin, "%d", &n)
	lookupByVal := map[int]int{}
	lookupByPos := map[int]int{}
	for i := 0; i < n; i++ {
		fmt.Fscanf(stdin, "%d", &x)
		lookupByVal[x] = i + 1
		lookupByPos[i+1] = x
	}
	for i := 0; i < n; i++ {
		pos := lookupByVal[i+1]
		pos2 := lookupByVal[pos]
		fmt.Fprintln(stdout, pos2)
	}
}
