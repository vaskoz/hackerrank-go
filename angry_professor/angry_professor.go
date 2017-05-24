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
	var tests int
	fmt.Fscanf(stdin, "%d", &tests)
	for t := 0; t < tests; t++ {
		var n, k, count, x int
		fmt.Fscanf(stdin, "%d %d", &n, &k)
		for i := 0; i < n; i++ {
			fmt.Fscanf(stdin, "%d", &x)
			if x <= 0 {
				count++
			}
		}
		if count < k {
			fmt.Fprintln(stdout, "YES")
		} else {
			fmt.Fprintln(stdout, "NO")
		}
	}
}
