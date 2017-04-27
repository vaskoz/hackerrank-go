package main

import (
	"fmt"
	"io"
	"os"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout

func main() {
	hackerrank := "hackerrank"
	var n int
	fmt.Fscanf(stdin, "%d", &n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscanf(stdin, "%s", &s)
		hrpos := 0
		for j := 0; j < len(s); j++ {
			if hackerrank[hrpos] == s[j] {
				hrpos++
			}
			if hrpos == len(hackerrank) {
				fmt.Fprintln(stdout, "YES")
				break
			}
		}
		if hrpos != len(hackerrank) {
			fmt.Fprintln(stdout, "NO")
		}
	}
}
