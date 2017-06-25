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
	var queries int
	fmt.Fscanf(stdin, "%d", &queries)
	for query := 0; query < queries; query++ {
		var n int
		fmt.Fscanf(stdin, "%d", &n)
		a, b := make([]int, n), make([]int, n)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				var x int
				fmt.Fscanf(stdin, "%d", &x)
				a[i] += x
				b[j] += x
			}
		}
		result := "Possible"
		for i := 0; i < n; i++ {
			j := 0
			for ; j < n; j++ {
				if a[i] == b[j] {
					a[i], b[j] = b[j], a[i]
					break
				}
			}
			if j == n {
				result = "Impossible"
				break
			}
		}
		fmt.Fprintln(stdout, result)
	}
}
