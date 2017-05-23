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
	var n, sock, result int
	fmt.Fscanf(stdin, "%d", &n)
	socks := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Fscanf(stdin, "%d", &sock)
		socks[sock]++
	}
	for _, c := range socks {
		result += (c / 2)
	}
	fmt.Fprintln(stdout, result)
}
