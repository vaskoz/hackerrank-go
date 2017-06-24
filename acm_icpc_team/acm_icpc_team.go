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
	var n, m int
	fmt.Fscanf(stdin, "%d %d", &n, &m)
	input := make([][]rune, n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscanf(stdin, "%s", &s)
		input[i] = []rune(s)
	}
	var topics, teams int
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			count := 0
			for k := 0; k < m; k++ {
				if input[i][k] == '1' || input[j][k] == '1' {
					count++
				}
			}
			if count > topics {
				topics = count
				teams = 1
			} else if count == topics {
				teams++
			}
		}
	}
	fmt.Fprintln(stdout, topics)
	fmt.Fprintln(stdout, teams)
}
