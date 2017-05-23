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
	a := make([]int, 26)
	for i := 0; i < len(a); i++ {
		fmt.Fscanf(stdin, "%d", &a[i])
	}
	var word string
	fmt.Fscanf(stdin, "%s", &word)
	height := 0
	for _, c := range word {
		if a[c-'a'] > height {
			height = a[c-'a']
		}
	}
	fmt.Fprintln(stdout, len(word)*height)
}
