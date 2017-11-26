package main

import (
	"fmt"
	"io"
	"os"
	"unicode"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout

func main() {
	var input string
	fmt.Fscanf(stdin, "%s", &input)
	result := 1
	for _, r := range input {
		if unicode.IsUpper(r) {
			result++
		}
	}
	fmt.Fprintln(stdout, result)
}
