/* *
 * Largest Product in a Grid
 * ----------------------
 * In the 20x20 grid below, four numbers along a diagonal line have been marked.

 * 				08 02 22 97 38 15 00 40  00  75  04  05  07 78 52 12 50 77 91 08
 *				49 49 99 40 17 81 18 57  60  87  17  40  98 43 69 48 04 56 62 00
 *				81 49 31 73 55 79 14 29  93  71  40  67  53 88 30 03 49 13 36 65
 *				52 70 95 23 04 60 11 42  69  24  68  56  01 32 56 71 37 02 36 91
 *				22 31 16 71 51 67 63 89  41  92  36  54  22 40 40 28 66 33 13 80
 *				24 47 32 60 99 03 45 02  44  75  33  53  78 36 84 20 35 17 12 50
 *				32 98 81 28 64 23 67 10 |26| 38  40  67  59 54 70 66 18 38 64 70
 *				67 26 20 68 02 62 12 20  95 |63| 94  39  63 08 40 91 66 49 94 21
 *				24 55 58 05 66 73 99 26  97  17 |78| 78  96 83 14 88 34 89 63 72
 *				21 36 23 09 75 00 76 44  20  45  35 |14| 00 61 33 97 34 31 33 95
 *				78 17 53 28 22 75 31 67  15  94  03  80  04 62 16 14 09 53 56 92
 *				16 39 05 42 96 35 31 47  55  58  88  24  00 17 54 24 36 29 85 57
 *				86 56 00 48 35 71 89 07  05  44  44  37  44 60 21 58 51 54 17 58
 *				19 80 81 68 05 94 47 69  28  73  92  13  86 52 17 77 04 89 55 40
 *				04 52 08 83 97 35 99 16  07  97  57  32  16 26 26 79 33 27 98 66
 *				88 36 68 87 57 62 20 72  03  46  33  67  46 55 12 32 63 93 53 69
 *				04 42 16 73 38 25 39 11  24  94  72  18  08 46 29 32 40 62 76 36
 *				20 69 36 41 72 30 23 88  34  62  99  69  82 67 59 85 74 04 36 16
 *				20 73 35 29 78 31 90 01  74  31  49  71  48 86 81 16 23 57 05 54
 *				01 70 54 71 83 51 54 69  16  92  33  48  61 43 52 01 89 19 67 48
 *
 * The product of these numbers is 26 * 64 * 78 * 14 = 1788693.
 *
 * What is the greatest product of four adjacent numbers in the same
 * direction (up, down, left, right, or diagonally) in the 20x20 grid?
 * ------------------------------------------------------------------
 */

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


func read_grid_from_file(file_path string) [][] uint64 {
	grid, err := os.ReadFile(file_path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSuffix(string(grid), "\n"), "\n")
	matrix := make([][] uint64, len(lines));

	for idx, line := range lines {
		numbers := strings.Split(line, " ")
		for _, num := range numbers {
			i, _ := strconv.Atoi(num)
			matrix[idx] = append(matrix[idx], uint64(i))
		}
	}
	return matrix
}

func calc_horizontal(matrix [][] uint64, rows int, cols int, row int, col int, n_numbers int) uint64 {
	if col + n_numbers >= cols{
		return 0;
	}

	res := uint64(1);
	for i := 0; i < n_numbers; i++ {
		res *= matrix[row][col+i];
	}

	return res;
}

func calc_vertical(matrix [][]uint64, rows int, cols int, row int, col int, n_numbers int) uint64 {
	if row + n_numbers >= rows{
		return 0;
	}

	res := uint64(1);
	for i := 0; i < n_numbers; i++ {
		res *= matrix[row+i][col];
	}

	return res;
}

func calc_diagonal_left_to_right(matrix [][]uint64, rows int, cols int, row int, col int, n_numbers int) uint64 {
	if row + n_numbers >= rows || col + n_numbers >= cols{
		return 0;
	}

	res := uint64(1);
	for i := 0; i < n_numbers; i++ {
		res *= matrix[row+i][col+i];
	}

	return res;
}

func calc_diagonal_right_to_left(matrix [][]uint64, rows int, cols int, row int, col int, n_numbers int) uint64 {
	if row + n_numbers >= rows || col + 1 - n_numbers < 0 {
		return 0;
	}
	res := uint64(1);
	for i := 0; i < n_numbers; i++ {
		res *= matrix[row+i][col-i];
	}

	return res;
}


func main() {
	grid := read_grid_from_file("./problem_011/grid.txt");
	rows := len(grid);
	cols := len(grid[0]);

	if calc_horizontal(grid, rows, cols, 0, 0, 4) != 34144  {
		panic("calc_horizontal(grid, rows, cols, 0, 0, 4) != 34144")
	}
	if calc_vertical(grid, rows, cols, 0, 0, 4) != 1651104  {
		panic("calc_vertical(grid, rows, cols, 0, 0, 4) != 1651104")
	}
	if calc_diagonal_left_to_right(grid, rows, cols, 0, 0, 4) != 279496  {
		panic("calc_diagonal_left_to_right(grid, rows, cols, 0, 0, 4) != 279496")
	}
	if calc_diagonal_right_to_left(grid, rows, cols, 0, 3, 4) != 24468444  {
		panic("calc_diagonal_right_to_left(grid, rows, cols, 0, 0, 4) != 24468444")
	}

	if calc_horizontal(grid, rows, cols, 0, 17, 4) != 0  {
		panic("calc_horizontal(grid, rows, rows-3, 0, 0, 4) != 0")
	}
	if calc_vertical(grid, rows, cols, 17, 0, 4) != 0  {
		panic("calc_vertical(grid, rows, cols, 0, cols-3, 4) != 0")
	}
	if calc_diagonal_left_to_right(grid, rows, cols, rows-3, cols-3, 4) != 0  {
		panic("calc_diagonal_left_to_right(grid, rows, cols, rows-3, cols-3, 4) != 0")
	}
	if calc_diagonal_right_to_left(grid, rows, cols, rows-3, 2, 4) != 0  {
		panic("calc_diagonal_right_to_left(grid, rows, cols, rows-3, 2, 4) != 0")
	}

	N_NUMBERS := 4
	max_result := uint64(0)
	max_result_row := 0
	max_result_col := 0
	max_result_dir := "horizontal"
	for row := 0; row < rows; row++ {
		for col:= 0; col < cols; col++ {
			result := calc_horizontal(grid, rows, cols, row, col, N_NUMBERS)
			if result > max_result {
				max_result = result
				max_result_row = row
				max_result_col = col
				max_result_dir = "horizontal"
			}
			result = calc_vertical(grid, rows, cols, row, col, N_NUMBERS)
			if result > max_result {
				max_result = result
				max_result_row = row
				max_result_col = col
				max_result_dir = "vertical"
			}
			result = calc_diagonal_left_to_right(grid, rows, cols, row, col, N_NUMBERS)
			if result > max_result {
				max_result = result
				max_result_row = row
				max_result_col = col
				max_result_dir = "left to right diagonal"
			}
			result = calc_diagonal_right_to_left(grid, rows, cols, row, col, N_NUMBERS)
			if result > max_result {
				max_result = result
				max_result_row = row
				max_result_col = col
				max_result_dir = "right to left diagonal"
			}
		}
	}
	fmt.Println("The greatest product of four adjacent numbers in the same direction: ", max_result, ".")
	fmt.Println("Position: (",max_result_col,", ", max_result_row,") in ", max_result_dir, "direction.")
}
