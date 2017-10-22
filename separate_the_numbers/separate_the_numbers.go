package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout

func main() {
	var q int
	fmt.Fscanf(stdin, "%d", &q)
	for i := 0; i < q; i++ {
		var s string
		fmt.Fscanf(stdin, "%s", &s)
		var found = false
		for l := 1; l <= len(s)/2; l++ {
			n, _ := strconv.Atoi(s[0:l])
			index := l
			for index < len(s) {
				n++
				x := fmt.Sprintf("%d", n)
				if index+len(x) > len(s) || s[index:index+len(x)] != x {
					break
				}
				index += len(x)
			}
			if index >= len(s) {
				found = true
				fmt.Fprintln(stdout, "YES", s[0:l])
				break
			}
		}
		if !found {
			fmt.Fprintln(stdout, "NO")
		}
	}
}
