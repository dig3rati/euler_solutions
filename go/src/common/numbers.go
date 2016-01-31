// Common Number related functions
// go 1.x
// @author: Bhaskar Nidumukalla
// @encoding: utf-8

package common

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

type PythogoreanTriplet struct {
	base, height uint64
	hypotenuse   uint64
}

// Generates a sequence of pythogorean triplet
// https://en.wikipedia.org/wiki/Formulas_for_generating_Pythagorean_triples
func PythogoreanTripletSeq() chan *PythogoreanTriplet {
	out := make(chan *PythogoreanTriplet)
	go func() {
		for i := uint64(1); true; i++ {
			pt := new(PythogoreanTriplet)
			pt.base = uint64(2*i + 1)
			pt.height = uint64(pt.base*i + i)
			pt.hypotenuse = uint64(math.Sqrt(float64(pt.base*pt.base + pt.height*pt.height)))
			out <- pt
		}
	}()
	return out
}

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

// Generates a continuous fibonacci sequence
func FibonacciContSeq() chan *big.Int {
	out := make(chan *big.Int)
	go func() {
		first, second := big.NewInt(1), big.NewInt(1)
		out <- first
		for f, s := first, second; true; f, s = s, f {
			out <- f
			s.Add(f, s)
		}
	}()
	return out
}

// Returns the circular numbers for a given number
func CircularNumbers(n uint64) []uint64 {
	bytes := []byte(fmt.Sprintf("%d", n))
	nums := make([]uint64, 0, len(bytes))
	nums = append(nums, n)
	for {
		var ld byte
		ld, bytes = bytes[0], bytes[1:]
		bytes = append(bytes, ld)
		cn, _ := strconv.ParseUint(string(bytes), 10, 64)
		if cn == n {
			break
		} else {
			nums = append(nums, cn)
		}
	}
	return nums
}

// Checks if the given number is a prime number
func IsPrime(number uint64) bool {
	return big.NewInt(int64(number)).ProbablyPrime(2)
}

// Checks if the given number is a circular prime
func IsCircularPrime(n uint64) bool {
	for _, i := range CircularNumbers(n) {
		if !IsPrime(i) {
			return false
		}
	}
	return true
}

// Returns the range between which nth prime number exists
// Based on https://en.wikipedia.org/wiki/Prime_number_theorem
func RangeOfNthPrime(n uint64) (uint64, uint64) {
	ln := math.Log(float64(n))
	lln := math.Log(ln)
	above := n * uint64(ln+lln-1)
	below := n * uint64(ln+lln)
	return above, below
}

// Returns an approximate number of primes below the given number
// Based on https://primes.utm.edu/howmany.html
func NumberOfPrimes(below uint64) uint64 {
	return uint64(float64(below) / (math.Log(float64(below)) - 1))
}

// Generates sequence of prime numbers below the given number
// Based on Sieve of Eratosthenes
// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
func PrimeSieveOfEratV1(upto uint64) chan uint64 {
	out := make(chan uint64)
	go func() {
		defer close(out)
		var slen uint64
		if upto > math.MaxInt32 {
			slen = uint64(math.MaxInt32)
		} else {
			slen = upto + 1
		}
		nots := make([]bool, slen)
		upto_sqrt := uint64(math.Sqrt(float64(upto)))
		for i := uint64(2); i <= upto_sqrt; i++ {
			for j, k := i*i, uint64(1); j <= upto; j, k = i*i+k*i, k+uint64(1) {
				if j > uint64(len(nots)) {
					nots = append(nots, make([]bool, upto-slen+1)...)
				}
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
	primes := PrimeSieveOfEratV1(number)
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
func Factorial(n *big.Int) *big.Int {
	var prod *big.Int = big.NewInt(1)
	for i := n; i.Cmp(big.NewInt(0)) > 0; i.Sub(i, big.NewInt(1)) {
		tmp := new(big.Int)
		tmp.Mul(prod, i)
		prod = tmp
	}
	return prod
}

// Returns the Power of x ^ y
func Pow(x *big.Int, y *big.Int) *big.Int {
	var pow *big.Int = big.NewInt(1)
	for i := big.NewInt(1); i.Cmp(y) <= 0; i.Add(i, big.NewInt(1)) {
		pow.Mul(pow, x)
	}
	return pow
}

// Returns the sum of self powers eg., 1^1 + 2^2 + 3^3 ... upto n ^ n
func SumOfSelfPowers(from *big.Int, to *big.Int) *big.Int {
	var sum *big.Int = big.NewInt(0)
	for i := from; i.Cmp(to) <= 0; i.Add(i, big.NewInt(1)) {
		sum.Add(sum, Pow(i, i))
	}
	return sum
}

// Returns number of digits in a number
func NumberOfDigits(n *big.Int) int {
	return len(n.String())
}
