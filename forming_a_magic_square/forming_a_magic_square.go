package main

import (
	"fmt"
	"io"
	"os"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr

var magic [][]int = [][]int{
	{8, 1, 6, 3, 5, 7, 4, 9, 2},
	{6, 1, 8, 7, 5, 3, 2, 9, 4},
	{4, 9, 2, 3, 5, 7, 8, 1, 6},
	{2, 9, 4, 7, 5, 3, 6, 1, 8},
	{8, 3, 4, 1, 5, 9, 6, 7, 2},
	{4, 3, 8, 9, 5, 1, 2, 7, 6},
	{6, 7, 2, 1, 5, 9, 8, 3, 4},
	{2, 7, 6, 9, 5, 1, 4, 3, 8},
}

func main() {
	matrix := make([]int, 9)
	for i := range matrix {
		fmt.Fscanf(stdin, "%d", &matrix[i])
	}
	result := 100
	for k := 0; k < len(magic); k++ {
		diff := 0
		for i := range matrix {
			if d := matrix[i] - magic[k][i]; d < 0 {
				diff += -d
			} else {
				diff += d
			}
		}
		if diff < result {
			result = diff
		}
	}
	fmt.Fprintln(stdout, result)
}
