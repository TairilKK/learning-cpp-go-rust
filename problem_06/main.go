/* *
 * Sum Square Difference
 * ----------------------
 * The sum of the squares of the first ten natural numbers is,
 * 				1^2 + 2^2 + ... + 10^2 = 385.
 *
 * The square of the sum of the first ten natural numbers is,
 * 			 (1 + 2 + ... + 10)^2 = 55^2 = 3025.
 *
 * Hence the difference between the sum of the squares of the
 * first ten natural numbers and the square of the sum is
 * 						3025 - 385 = 2640
 *
 * Find the difference between the sum of the squares of the first
 * one hundred natural numbers and the square of the sum.
 * ------------------------------------------------------------------
 */

package main

import "fmt"

func sum_of_squares_n_first(n uint64) uint64{
	var result uint64 = 0;
	for i:=uint64(1); i<=n; i++ {
		result += i*i;
	}
	return result;
}

func square_of_sum_n_first(n uint64) uint64{
	var result uint64 = 0;
	for i:=uint64(1); i<=n; i++ {
		result += i;
	}
	return result * result;
}

func main() {
	var test uint64 = 10
    fmt.Println(square_of_sum_n_first(test), sum_of_squares_n_first(test), square_of_sum_n_first(test) - sum_of_squares_n_first(test))


	var value uint64 = 100
    fmt.Println(square_of_sum_n_first(value), sum_of_squares_n_first(value), square_of_sum_n_first(value) - sum_of_squares_n_first(value))
}
