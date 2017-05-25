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
	var n, k int
	fmt.Fscanf(stdin, "%d %d", &n, &k)
	clouds := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(stdin, "%d", &clouds[i])
	}

	e := 100
	pos := 0
	for {
		pos = (pos + k) % n
		if clouds[pos] == 0 {
			e = e - 1
		} else {
			e = e - 3
		}
		if pos == 0 {
			break
		}
	}
	fmt.Fprintln(stdout, e)
}
