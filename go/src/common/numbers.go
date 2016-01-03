// Common Number related functions
// go 1.x
// @author: Bhaskar Nidumukalla
// @encoding: utf-8

package common

import (
	"math"
	"math/big"
)

// Generates the fibonacci sequence below the `max` number while sending the result to channel
func FibonacciSeq(start uint64, second uint64, max uint64) chan uint64 {
	out := make(chan uint64)
	go func() {
		out <- start
		defer close(out)
		for f, s := start, second; s <= max; f, s = s, f+s {
			out <- s
		}
	}()
	return out
}

// Checks if the given number is a prime number
func IsPrime(number uint64) bool {
	return big.NewInt(int64(number)).ProbablyPrime(2)
}

// Generates sequence of prime numbers below the given number
// Based on Sieve of Eratosthenes
// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
func PrimeSeq(upto uint64) chan uint64 {
	out := make(chan uint64)
	go func() {
		defer close(out)
		nots := make([]bool, upto+1)
		upto_sqrt := uint64(math.Sqrt(float64(upto)))
		for i := uint64(2); i <= upto_sqrt; i++ {
			for j, k := i*i, uint64(1); j <= upto; j, k = i*i+k*i, k+uint64(1) {
				nots[j] = true
			}
		}
		for i, n := range nots {
			if i == 0 || i == 1 {
				continue
			}
			if !n {
				out <- uint64(i)
			}
		}
	}()
	return out
}

// Returns the prime factors for a given number
func PrimeFactors(number uint64) chan uint64 {
	primes := PrimeSeq(number)
	out := make(chan uint64)
	go func() {
		defer close(out)
		for p := range primes {
			if number%p == 0 {
				out <- p
			}
		}
	}()
	return out
}

// GCD of given numbers
func GCD(a, b uint64) uint64 {
	var tmp uint64
	for b != 0 {
		tmp = a % b
		a = b
		b = tmp
	}
	return a
}

// LCM of given numbers
func LCM(a, b uint64) uint64 {
	mul := a * b
	var tmp uint64
	for b != 0 {
		tmp = a % b
		a = b
		b = tmp
	}
	return mul / a
}

// Sum of sequence of number from 1 to n
func SumOfNaturalNumbers(n uint64) uint64 {
	return n * (n + 1) / 2
}

// Sum of sequence of squares of numbers from 1 to n
func SumOfNaturalSquares(n uint64) uint64 {
	return n * (n + 1) * (2*n + 1) / 6
}

// Returns the factorial of a given number
func Factorial(n uint64) uint64 {
	var prod uint64 = 1
	for i := n; i > 0; i-- {
		prod *= i
	}
	return prod
}
