package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr

func main() {
	var n int
	fmt.Fscanf(stdin, "%d", &n)
	for i := 1; i <= n; i++ {
		fmt.Fprintf(stdout, "%s%s\n", strings.Repeat(" ", n-i), strings.Repeat("#", i))
	}
}
