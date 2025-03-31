/* *
 * 10 001st Prime
 * ----------------------
 * By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
 * we can see that the 6th prime is 13.
 *
 * What is the 10 001st prime number?
 * ------------------------------------------------------------------
 */

fn is_prime(n: u32) -> bool {
    if n <= 1 {
        return false;
    }
    let n_square = n.isqrt();
    for i in 2..n_square+1 {
        if n % i == 0 {
            return false
        }
    }
    return true;
}

fn main() {
    let mut counter = 1;
    let mut value = 1;
    while counter < 10002 {
        if is_prime(value) {
            println!("{}. {}", counter, value);
            counter += 1;
        }
        value += 1;
    }
}
