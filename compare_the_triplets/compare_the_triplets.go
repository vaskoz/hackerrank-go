package main

import (
	"fmt"
	"io"
	"os"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout

func main() {
	a := make([]int, 3)
	b := make([]int, 3)
	for i := range a {
		fmt.Fscanf(stdin, "%d", &a[i])
	}
	for i := range b {
		fmt.Fscanf(stdin, "%d", &b[i])
	}
	aliceScore, bobScore := 0, 0
	for i := range a {
		if a[i] > b[i] {
			aliceScore++
		} else if a[i] < b[i] {
			bobScore++
		}
	}
	fmt.Fprintln(stdout, aliceScore, bobScore)
}
