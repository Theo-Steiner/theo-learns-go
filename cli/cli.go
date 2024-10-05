package main

import (
	"fmt"
	"os"
)

// This program reads numbers from stdin and calculates the average.
// Try it out by writing a heredoc into the program's stdin like so:
// go run cli.go << EOF
// 1 3
// 2 4
// 3 5
// EOF
func main() {
	var n int
	var sum float64

	for {
		var val float64
		// try to scan the next content from stdin as a float and write it to memory at the pointer val
		_, err := fmt.Fscanln(os.Stdin, &val)
		if err != nil {
			break
		}

		sum += val
		n++
	}

	if n == 0 {
		fmt.Fprintln(os.Stderr, "No numbers provided")
		os.Exit(-1)
	}

	fmt.Fprintln(os.Stdout, "The average is", sum/float64(n))
}
