/* *
 * Factorial Digit Sum
 * ----------------------
 * n! means n times (n - 1) times cdots times 3 times 2 times 1.
 * For example,
 *              10! = 10 * 9 * ... * 3 * 2 * 1 = 3628800,
 * and the sum of the digits in the number 10! is
 *              3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
 *
 * Find the sum of the digits in the number 100!.
 */

 use num_bigint::BigUint;
 use num_traits::One;

 fn digit_sum(number: &BigUint) -> u32 {
    number.to_string()
          .chars()
          .map(|c| c.to_digit(10).unwrap())
          .sum()
}

 fn factorial_digit_sum(n: u64) -> BigUint {
    let mut result: u64 = BigUint::one();
    for i in 1..n+1 {
        result *= i;
        if result % 10 == 0 {
            result /= 10;
        }
    }
    digit_sum(result)
 }

fn main() {
    let n: u64 = 10;
    assert!(factorial_digit_sum(n) == 27);
    println!("The sum of the digits in the number 10! is {}", factorial_digit_sum(10));
    println!("The sum of the digits in the number 100! is {}", factorial_digit_sum(100));
}
