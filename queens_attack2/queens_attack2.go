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
	for r, c := qr, qc+1; r <= n && r > 0 && c <= n && c > 0; c, result = c+1, result+1 {
		if _, found := obstacles[pos{r, c}]; found {
			break
		}
	}
	for r, c := qr, qc-1; r <= n && r > 0 && c <= n && c > 0; c, result = c-1, result+1 {
		if _, found := obstacles[pos{r, c}]; found {
			break
		}
	}
	for r, c := qr+1, qc; r <= n && r > 0 && c <= n && c > 0; r, result = r+1, result+1 {
		if _, found := obstacles[pos{r, c}]; found {
			break
		}
	}
	for r, c := qr-1, qc; r <= n && r > 0 && c <= n && c > 0; r, result = r-1, result+1 {
		if _, found := obstacles[pos{r, c}]; found {
			break
		}
	}
	for r, c := qr+1, qc+1; r <= n && r > 0 && c <= n && c > 0; r, c, result = r+1, c+1, result+1 {
		if _, found := obstacles[pos{r, c}]; found {
			break
		}
	}
	for r, c := qr-1, qc-1; r <= n && r > 0 && c <= n && c > 0; r, c, result = r-1, c-1, result+1 {
		if _, found := obstacles[pos{r, c}]; found {
			break
		}
	}
	for r, c := qr+1, qc-1; r <= n && r > 0 && c <= n && c > 0; r, c, result = r+1, c-1, result+1 {
		if _, found := obstacles[pos{r, c}]; found {
			break
		}
	}
	for r, c := qr-1, qc+1; r <= n && r > 0 && c <= n && c > 0; r, c, result = r-1, c+1, result+1 {
		if _, found := obstacles[pos{r, c}]; found {
			break
		}
	}
	fmt.Fprintln(stdout, result)
}
