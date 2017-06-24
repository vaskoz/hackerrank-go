package main

import (
	"fmt"
	"io"
	"os"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr

type pos struct {
	r, c int
}

func main() {
	var n, k, qr, qc, r, c, result int
	fmt.Fscanf(stdin, "%d %d", &n, &k)
	fmt.Fscanf(stdin, "%d %d", &qr, &qc)
	obstacles := make(map[pos]struct{})
	for i := 0; i < k; i++ {
		fmt.Fscanf(stdin, "%d %d", &r, &c)
		obstacles[pos{r, c}] = struct{}{}
	}
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}}
	for _, d := range directions {
		for r, c := qr, qc; r <= n && r > 0 && c <= n && c > 0; r, c, result = r+d[0], c+d[1], result+1 {
			if r == qr && c == qc {
				result--
			} else if _, found := obstacles[pos{r, c}]; found {
				break
			}
		}
	}
	fmt.Fprintln(stdout, result)
}
