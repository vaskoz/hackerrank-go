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
	var games, best, worst, worstBroken, bestBroken int
	fmt.Fscanf(stdin, "%d", &games)
	fmt.Fscanf(stdin, "%d", &best)
	worst = best
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
