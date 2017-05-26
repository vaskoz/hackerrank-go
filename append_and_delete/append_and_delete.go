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
	var s, t string
	var k int
	fmt.Fscanf(stdin, "%s", &s)
	fmt.Fscanf(stdin, "%s", &t)
	fmt.Fscanf(stdin, "%d", &k)

	index := 0
	for i := 0; i < len(s) && i < len(t); i++ {
		if s[i] != t[i] {
			break
		} else {
			index++
		}
	}
	if combined := len(s) + len(t); combined-2*index > k {
		fmt.Fprintln(stdout, "No")
	} else if (combined-2*index)%2 == k%2 {
		fmt.Fprintln(stdout, "Yes")
	} else if combined-k < 0 {
		fmt.Fprintln(stdout, "Yes")
	} else {
		fmt.Fprintln(stdout, "No")
	}
}
