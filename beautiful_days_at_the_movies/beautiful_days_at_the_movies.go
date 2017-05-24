package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr

func main() {
	var i, j, k, counter int
	fmt.Fscanf(stdin, "%d %d %d", &i, &j, &k)
	for x := i; x <= j; x++ {
		r := Reverse(x)
		v := x - r
		if v%k == 0 {
			counter++
		}
	}
	fmt.Fprintln(stdout, counter)
}

func Reverse(x int) int {
	return ReverseUsingInt(x)
	//return ReverseUsingString(x)
}

func ReverseUsingString(x int) int {
	s := strconv.Itoa(x)
	letters := strings.Split(s, "")
	for i := 0; i < len(letters)/2; i++ {
		letters[i], letters[len(letters)-i-1] = letters[len(letters)-i-1], letters[i]
	}
	s = strings.Join(letters, "")
	result, _ := strconv.Atoi(s)
	return result
}

func ReverseUsingInt(x int) int {
	result := 0
	for x/10 != 0 {
		if x%10 != 0 {
			result += (x % 10)
		}
		result *= 10
		x /= 10
	}
	result += x
	return result
}
