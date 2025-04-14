/* *
 * Longest Collatz Sequence
 * ----------------------
 * The following iterative sequence is defined for the set of positive integers:</p>
 * 								n -> n/2 (n is even)
 * Using the rule above and starting with $13$, we generate the following sequence:
 * 								n -> 3n + 1 (n is odd)
 * 13 -> 40 -> 20 -> 10 -> 5 -> 16 -> 8 -> 4 -> 2 -> 1.
 * It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
 * Which starting number, under one million, produces the longest chain?
 *
 * NOTE: Once the chain starts the terms are allowed to go above one million.
 * ------------------------------------------------------------------
 */

package main

import "fmt"

func colatz_sequence(n uint64) []uint64 {
	var sequence []uint64
	sequence = append(sequence, n)
	for n != 1 {
		if n % 2 == 0 {
			n /= 2
		} else {
			n = 3*n+1
		}
		sequence = append(sequence, n)
	}

	return sequence
}

func main() {
	var longest_sequence []uint64
	for n:=uint64(1); n < 1_000_000; n++ {
		current_sequence := colatz_sequence(n)
		if len(longest_sequence) < len(current_sequence) {
			longest_sequence = current_sequence
		}
	}

	fmt.Println("Longest Collatz Sequence with starting number, under one million starts with", longest_sequence[0])
	for idx, value := range(longest_sequence) {
		fmt.Print(value, " ");
		if idx % 10 == 9{
			fmt.Print("\n");
		}
	}
}
