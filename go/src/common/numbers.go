// Common Number related functions
// go 1.x
// @author: Bhaskar Nidumukalla
// @encoding: utf-8

package common

// Generates the fibonacci sequence below the `stop` number while sending the result to `out` channel
func Fibonacci(start uint32, second uint32, max uint32, out chan<- uint32) {
	out <- start
	defer close(out)
	for f, s := start, second; s < max; f, s = s, f+s {
		out <- s
	}
}
