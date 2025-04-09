/* *
 * Number Letter Counts
 * ----------------------
 * If the numbers 1 to 5 are written out in words: one, two, three, four, five,
 * then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
 *
 * If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words,
 * how many letters would be used?
 *
 * NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two)
 * contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of
 * "and" when writing out numbers is in compliance with British usage.
 * ------------------------------------------------------------------
 */

 const NUMBERS: [&str; 20] = [
    "zero",
    "one",
    "two",
    "three",
    "four",
    "five",
    "six",
    "seven",
    "eight",
    "nine",
    "ten",
    "eleven",
    "twelve",
    "thirteen",
    "fourteen",
    "fifteen",
    "sixteen",
    "seventeen",
    "eighteen",
    "nineteen",
];

const TENS: [&str; 8] = [
    "twenty",
    "thirty",
    "forty",
    "fifty",
    "sixty",
    "seventy",
    "eighty",
    "ninety"
];

fn number_to_letters(n: u32) -> String {
    match n {
        0..=19 => NUMBERS[n as usize].to_string(),
        20..=99 => {
            let tens = TENS[(n / 10 - 2) as usize];
            let units = n % 10;
            if units == 0 {
                tens.to_string()
            } else {
                format!("{}-{}", tens, NUMBERS[units as usize])
            }
        },
        100..=999 => {
            let hundreds = NUMBERS[(n / 100) as usize];
            let remainder = n % 100;
            if remainder == 0 {
                format!("{} hundred", hundreds)
            } else {
                format!("{} hundred and {}", hundreds, number_to_letters(remainder))
            }
        },
        1000 => "one thousand".to_string(),
        _ => panic!("Number out of range"),
    }
}

fn main() {
    let mut total_letters = 0;
    for i in 1..=1000 {  // Zaczynamy od 1, bo 0 ("zero") nie jest liczone w zadaniu
        let word = number_to_letters(i);
        let letters = word.replace(" ", "").replace("-", "").len();
        total_letters += letters;
    }
    println!("Total letters used: {}", total_letters);
}
