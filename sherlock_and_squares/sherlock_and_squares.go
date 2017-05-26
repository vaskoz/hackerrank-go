package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr

func main() {
	var tests, a, b int
	fmt.Fscanf(stdin, "%d", &tests)
	for t := 0; t < tests; t++ {
		fmt.Fscanf(stdin, "%d %d", &a, &b)
		low := int(math.Sqrt(float64(a)))
		count := 0
		for i := low; math.Pow(float64(i), float64(2)) <= float64(b); i++ {
			if math.Pow(float64(i), float64(2)) >= float64(a) {
				count++
			}
		}
		fmt.Fprintln(stdout, count)
	}
}
