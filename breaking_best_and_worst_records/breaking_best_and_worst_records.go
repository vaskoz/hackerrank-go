package main

import (
	"fmt"
	"io"
	"os"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout

func main() {
	var games int
	fmt.Fscanf(stdin, "%d", &games)
	var best, worst int
	fmt.Fscanf(stdin, "%d", &best)
	worst = best
	var worstBroken, bestBroken int
	for i := 0; i < games-1; i++ {
		var score int
		fmt.Fscanf(stdin, "%d", &score)
		if score < worst {
			worstBroken++
			worst = score
		} else if score > best {
			bestBroken++
			best = score
		}
	}
	fmt.Fprintln(stdout, bestBroken, worstBroken)
}
