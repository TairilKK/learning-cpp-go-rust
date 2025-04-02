/* *
 * Special Pythagorean Triplet
 * ----------------------
 * A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
 * 							a^2 + b^2 = c^2.
 *
 * For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
 *
 * There exists exactly one Pythagorean triplet for which a + b + c = 1000.
 * Find the product abc.
 * ------------------------------------------------------------------
 */

package main

import "fmt"

func is_pythagorean_triplet(a uint32, b uint32, c uint32) bool {
	return (a*a + b*b) == (c*c);
}

func find_triplet(equals uint32) (uint32, uint32, uint32){
	var a, b, c uint32 = 1, 2, 3;
	for c < equals {
		b = 1;
		for b < c {
			a = 1;
			for a < b {
				if is_pythagorean_triplet(a, b, c) && a + b + c == equals{
					return a, b, c;
				}
				a++;
			}
			b++;
		}
		c++;
	}

	return 0, 0, 0;
}

func main() {
	var a, b, c = find_triplet(3+4+5);
	fmt.Println("Triplet:");
	fmt.Println(a, " + ", b, " + ", c, " = ", a+b+c);
	fmt.Println("Triplet product:");
	fmt.Println(a, " * ", b, " * ", c, " = ", a*b*c);
	if a != 3 && b != 4 && c != 5{
		panic("Incorrect values")
	}

	a, b, c = find_triplet(1000);
	fmt.Println("Triplet:");
	fmt.Println(a, " + ", b, " + ", c, " = ", a+b+c);
	fmt.Println("Triplet product:");
	fmt.Println(a, " * ", b, " * ", c, " = ", a*b*c);
}
