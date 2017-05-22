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
	var n, k, count int
	fmt.Fscanf(stdin, "%d %d", &n, &k)
	data := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(stdin, "%d", &data[i])
	}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if pair := data[i] + data[j]; pair/k > 0 && pair%k == 0 {
				count++
			}
		}
	}
	fmt.Fprintln(stdout, count)
}
