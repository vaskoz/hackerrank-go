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
	var s, n, m int
	fmt.Fscanf(stdin, "%d %d %d", &s, &n, &m)
	keyboards := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(stdin, "%d", &keyboards[i])
	}
	usbs := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscanf(stdin, "%d", &usbs[i])
	}
	sort.Ints(keyboards)
	sort.Ints(usbs)
	if keyboards[0]+usbs[0] > s {
		fmt.Fprintln(stdout, -1)
		return
	}
	kimax := sort.Search(len(keyboards), func(i int) bool {
		return keyboards[i] >= s
	})
	uimax := sort.Search(len(usbs), func(i int) bool {
		return usbs[i] >= s
	})
	max := 0
	for i := 0; i < kimax; i++ {
		for j := 0; j < uimax; j++ {
			if total := keyboards[i] + usbs[j]; total <= s && total > max {
				max = total
			}
		}
	}
	fmt.Fprintln(stdout, max)
}
