/* *Multiples of 3 and 5*
 * If we list all the natural numbers below 10 that are multiples of 3 or 5,
 * we get 3, 5, 6 and 9. The sum of these multiples is 23.
 * Find the sum of all the multiples of 3 or 5 below 1000.
 *
 * https://projecteuler.net/problem=1
 */

package problem_1

func SumOfMultiples(multiples []uint64, upto uint64) uint64 {
	var sum uint64
	for num := uint64(1); num < upto; num++ {
		for _, m := range multiples {
			if num%m == 0 {
				sum += num
				break
			}
		}
	}
	return sum
}
