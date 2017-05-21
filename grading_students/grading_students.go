package main

import (
	"fmt"
	"io"
	"os"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr

func main() {
	var grades int
	fmt.Fscanf(stdin, "%d", &grades)
	for i := 0; i < grades; i++ {
		var grade int
		fmt.Fscanf(stdin, "%d", &grade)
		if grade < 38 || grade%5 < 3 {
			fmt.Fprintln(stdout, grade)
		} else {
			fmt.Fprintln(stdout, (grade/5+1)*5)
		}
	}
}
