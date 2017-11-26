package main

import (
	"fmt"
	"io"
	"os"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout

func main() {
	var k int
	var unencrypted, encrypted string
	fmt.Fscanf(stdin, "%d", &k) // throw away length don't need it
	fmt.Fscanf(stdin, "%s", &unencrypted)
	fmt.Fscanf(stdin, "%d", &k)
	for _, r := range unencrypted {
		if r >= 'A' && r <= 'Z' {
			r += rune(k % 26)
			if r > 'Z' {
				r -= rune(26)
			}
		} else if r >= 'a' && r <= 'z' {
			r += rune(k % 26)
			if r > 'z' {
				r -= rune(26)
			}
		}
		encrypted += string(r)
	}
	fmt.Fprintln(stdout, encrypted)
}
