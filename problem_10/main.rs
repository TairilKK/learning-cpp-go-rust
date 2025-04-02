/* *
 * Summation of Primes
 * ----------------------
 * The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
 *
 * Find the sum of all the primes below two million.
 * ------------------------------------------------------------------
 */

fn is_prime(n: u64) -> u64 {
    let n_sqrt: u64 = n.isqrt();
    if n < 2 {
        return 0;
    }
    for i in 2..n_sqrt+1 {
        if n % i == 0 {
            return 0;
        }
    }
    return n;
}

fn main() {
    let mut sum: u64 = 0;
    for i in 1..10 {
        sum += is_prime(i);
    }
    assert!(sum == 17, "Sum should equals 17, but it is {}", sum);
    println!("Prime numbers sum below 10 is: {}", sum);

    let mut sum: u64 = 0;
    for i in 1..2000000 {
        sum += is_prime(i);
    }
    println!("Prime numbers sum below 2 000 000 is: {}", sum);
}
