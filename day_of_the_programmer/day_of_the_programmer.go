package main

import (
	"fmt"
	"io"
	"os"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout

func main() {
	var year int
	fmt.Fscanf(stdin, "%d", &year)
	month := 9
	days := 13
	if year > 1918 {
		if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
			days--
		}
	} else if year == 1918 {
		days = 26
	} else {
		if year%4 == 0 {
			days--
		}
	}
	fmt.Fprintf(stdout, "%02d.%02d.%4d\n", days, month, year)
}
