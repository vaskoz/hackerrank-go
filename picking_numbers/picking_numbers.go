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
	m := map[int]int{}
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscanf(stdin, "%d", &x)
		m[x]++
	}
	type tuple struct{ freq, num int }
	var tuples []tuple
	for num, freq := range m {
		tuples = append(tuples, tuple{num: num, freq: freq})
	}
	sort.Slice(tuples, func(i, j int) bool {
		return tuples[i].num < tuples[j].num
	})
	max := tuples[0].freq
	for i := 1; i < len(tuples); i++ {
		if freq := tuples[i].freq; freq > max {
			max = freq
		}
		if diff := tuples[i-1].num - tuples[i].num; diff >= -1 && diff <= 1 {
			sum := tuples[i-1].freq + tuples[i].freq
			if sum > max {
				max = sum
			}
		}
	}
	fmt.Fprintln(stdout, max)
}
