package main

import (
	"fmt"
	"io"
	"os"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout

func main() {
	var s string
	fmt.Fscanf(stdin, "%s", &s)
	u := make(map[int]struct{})
	var last byte = 0
	freq := 1
	for i := 0; i < len(s); i++ {
		if val := int(s[i] - 'a' + 1); s[i] != last {
			last = s[i]
			freq = 1
			u[val] = struct{}{}
		} else {
			freq++
			u[val*freq] = struct{}{}
		}
	}
	var n int
	fmt.Fscanf(stdin, "%d", &n)
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscanf(stdin, "%d", &x)
		if _, found := u[x]; found {
			fmt.Fprintln(stdout, "Yes")
		} else {
			fmt.Fprintln(stdout, "No")
		}
	}
}
