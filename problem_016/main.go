/* *
 * Power Digit Sum
 * ----------------------
 * 2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26
 *
 * What is the sum of the number 2^1000
 * ------------------------------------------------------------------
 */

package main

import (
	"fmt"
	"math/big"
)

func digits_sum(n *big.Int) *big.Int {
	result := new(big.Int)
	dig := new(big.Int)
	for _, digit := range(n.String()) {
		dig.SetString(string(digit), 10)
		result.Add(result, dig)
	}
	return result
}

func main() {
	test_res := big.NewInt(26)
	test := big.NewInt(2)
	test = test.Exp(big.NewInt(2), big.NewInt(15), nil)
	if test_res.Cmp(digits_sum(test)) != 0 {
		message := fmt.Sprintf("Result shoul equal to %s but it is %s.", test_res.String(), digits_sum(test))
		panic(message)
	}

	base := big.NewInt(2)
	exp := big.NewInt(1000)
	value := big.NewInt(0)
	value = value.Exp(base, exp, nil)
	dig_sum := digits_sum(value)
	message := fmt.Sprintf("%s^%s = %s and the sum of its digits is %s.", base.String(), exp.String(), value.String(), dig_sum.String())
	fmt.Println(message)

}
