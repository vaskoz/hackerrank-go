package main

import (
	"fmt"
	"io"
	"os"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout

func main() {
	var n string
	var k int
	fmt.Fscanf(stdin, "%s %d", &n, &k)
	sum := 0
	for _, r := range n {
		sum += int(r - '0')
	}
	fmt.Fprintln(stdout, DigitSum(fmt.Sprint(sum*k)))
}

func DigitSum(x string) string {
	if len(x) == 1 {
		return x
	}
	sum := 0
	for _, r := range x {
		sum += int(r - '0')
	}
	x = fmt.Sprint(sum)
	if len(x) > 1 {
		x = DigitSum(x)
	}
	return x
}
