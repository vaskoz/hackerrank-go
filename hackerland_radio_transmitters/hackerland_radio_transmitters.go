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
	var n, k int
	fmt.Fscanf(stdin, "%d %d", &n, &k)
	houses := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(stdin, "%d", &houses[i])
	}
	sort.Ints(houses)
	antennas := 0
	pos := houses[0]
	sum := 0
	for i := 1; i < n; i++ {
		if diff := houses[i] - pos; diff <= k {
			sum += diff
			if sum <= k {
				pos = houses[i]
			}
		} else {
			pos = houses[i]
			antennas++
			sum = 0
		}
	}
	fmt.Fprintln(stdout, antennas+1)
}
