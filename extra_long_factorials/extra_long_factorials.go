package main

import (
	"fmt"
	"io"
	"math/big"
	"os"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr

func main() {
	var N int
	fmt.Fscanf(stdin, "%d", &N)
	fmt.Fprintln(stdout, factorial(N).String())
}

func factorial(n int) *big.Int {
	result, _ := new(big.Int).SetString("1", 10)
	for i := 2; i <= n; i++ {
		x := new(big.Int).SetInt64(int64(i))
		result.Mul(result, x)
	}
	return result
}
