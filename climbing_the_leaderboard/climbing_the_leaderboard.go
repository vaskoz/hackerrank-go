package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout

func main() {
	var n int
	fmt.Fscanf(stdin, "%d", &n)
	var scores []int
	for i := 0; i < n; i++ {
		var s int
		fmt.Fscanf(stdin, "%d", &s)
		if scores == nil || scores[len(scores)-1] != s {
			scores = append(scores, s)
		}
	}
	var m int
	fmt.Fscanf(stdin, "%d", &m)
	for i := 0; i < m; i++ {
		var alice int
		fmt.Fscanf(stdin, "%d", &alice)
		index := sort.Search(len(scores), func(i int) bool {
			return scores[i] <= alice
		})
		fmt.Fprintln(stdout, index+1)
	}
}
