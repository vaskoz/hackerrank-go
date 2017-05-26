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
	var am, ad, ay int
	fmt.Fscanf(stdin, "%d %d %d", &ad, &am, &ay)
	var em, ed, ey int
	fmt.Fscanf(stdin, "%d %d %d", &ed, &em, &ey)

	fine := 0
	if ay > ey {
		fine = 10000
	} else if ay == ey && am > em {
		fine = 500 * (am - em)
	} else if ay == ey && am == em && ad > ed {
		fine = 15 * (ad - ed)
	}
	fmt.Fprintln(stdout, fine)
}
