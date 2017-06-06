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
	var s string
	var n int
	fmt.Fscanf(stdin, "%s", &s)
	fmt.Fscanf(stdin, "%d", &n)
	full := n / len(s)
	rem := n % len(s)
	var aCount, aPartial int
	for i, r := range s {
		if r == 'a' {
			aCount++
			if i < rem {
				aPartial++
			}
		}
	}
	result := full*aCount + aPartial
	fmt.Fprintln(stdout, result)
}
