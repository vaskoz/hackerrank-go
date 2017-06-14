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
	var line string
	for i := 0; i < tests; i++ {
		fmt.Fscanf(stdin, "%s", &line)
		if len(line) < 2 {
			fmt.Fprintln(stdout, "no answer")
			continue
		}
		runes := []rune(line)
		pivotIndex := -1
		for j := len(runes) - 1; j > 0; j-- {
			if runes[j] > runes[j-1] {
				pivotIndex = j - 1
				break
			}
		}
		if pivotIndex == -1 {
			fmt.Fprintln(stdout, "no answer")
			continue
		}
		nextIndex := -1
		for j := len(runes) - 1; j > pivotIndex; j-- {
			if runes[j] > runes[pivotIndex] {
				nextIndex = j
				break
			}
		}
		runes[pivotIndex], runes[nextIndex] = runes[nextIndex], runes[pivotIndex]
		for x, y := pivotIndex+1, len(runes)-1; x < y; x, y = x+1, y-1 {
			runes[x], runes[y] = runes[y], runes[x]
		}
		fmt.Fprintln(stdout, string(runes))
	}
}
