/* *
 * Large Sum
 * ----------------------
 * Work out the first ten digits of the sum of the following one-hundred
 * 50-digit numbers
 * ------------------------------------------------------------------
 */

package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"
)


func read_grid_from_file(file_path string) [] *big.Int {
	grid, err := os.ReadFile(file_path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSuffix(string(grid), "\n"), "\n")
	matrix := make([] *big.Int, len(lines));

	for idx, line := range lines {
		matrix[idx] = new(big.Int)
		matrix[idx].SetString(line, 10)
	}
	return matrix
}

func main() {

	numbers := read_grid_from_file("./problem_013/100_50-digits_numbers.txt");

	sum := new(big.Int)
	sum.SetString("0", 10)
	for _, number := range(numbers) {
		sum.Add(sum, number)
	}

	fmt.Println("Sum of this numbers is equal to", sum);
	fmt.Println("First 10-digits:", sum.String()[:10]);
}
