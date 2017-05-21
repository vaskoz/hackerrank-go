package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr

func main() {
	var line string
	fmt.Fscanln(stdin, &line)
	if meridiem := line[len(line)-2:]; meridiem == "AM" {
		if line[0:2] == "12" {
			fmt.Fprintf(stdout, "00%s\n", line[2:len(line)-2])
		} else {
			fmt.Fprintln(stdout, line[0:len(line)-2])
		}
	} else {
		if line[0:2] == "12" {
			fmt.Fprintf(stdout, "12%s\n", line[2:len(line)-2])
		} else {
			h, _ := strconv.Atoi(line[0:2])
			fmt.Fprintf(stdout, "%d%s\n", h+12, line[2:len(line)-2])
		}
	}
}
