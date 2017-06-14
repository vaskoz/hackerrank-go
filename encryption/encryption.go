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
	var msg string
	fmt.Fscanf(stdin, "%s", &msg)
	L := len(msg)
	row := int(math.Sqrt(float64(L)))
	col := row
	for row*col < L {
		if col > row {
			row++
		} else if row == col {
			col++
		}
	}
	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			p := col*j + i
			if p < L {
				fmt.Fprintf(stdout, "%s", string(msg[p]))
			}
		}
		fmt.Fprint(stdout, " ")
	}
}
