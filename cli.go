package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	var sum float64
	var secondSum float64

	for {
		var val float64
		var second float64
		_, err := fmt.Fscanln(os.Stdin, &val, &second)
		if err != nil {
			break
		}

		sum += val
		secondSum += second
		n++
	}

	if n == 0 {
		fmt.Fprintln(os.Stderr, "No numbers provided")
		os.Exit(-1)
	}

	fmt.Fprintln(os.Stdout, "The average is", sum/float64(n))
	fmt.Fprintln(os.Stdout, "The second average is", secondSum/float64(n))
}
