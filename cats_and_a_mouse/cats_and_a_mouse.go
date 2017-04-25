package main

import (
	"fmt"
	"io"
	"os"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout

func main() {
	var q int
	fmt.Fscanf(stdin, "%d", &q)
	for i := 0; i < q; i++ {
		var x, y, z int
		fmt.Fscanf(stdin, "%d %d %d", &x, &y, &z)
		da := z - x
		db := z - y
		if da < 0 {
			da = -da
		}
		if db < 0 {
			db = -db
		}
		if da < db {
			fmt.Fprintln(stdout, "Cat A")
		} else if db < da {
			fmt.Fprintln(stdout, "Cat B")
		} else {
			fmt.Fprintln(stdout, "Mouse C")
		}
	}
}
