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
	var n, k, x, result int
	fmt.Fscanf(stdin, "%d %d", &n, &k)
	mod := make([]int, k)
	for i := 0; i < n; i++ {
		fmt.Fscanf(stdin, "%d", &x)
		mod[x%k]++ // store count of remainders
	}
	for i := 0; i < k/2+1; i++ {
		if i == 0 || k == i*2 {
			if mod[i] != 0 {
				result += 1
			}
		} else {
			if mod[i] > mod[k-i] {
				result += mod[i]
			} else {
				result += mod[k-i]
			}
		}
	}
	fmt.Fprintln(stdout, result)
}
