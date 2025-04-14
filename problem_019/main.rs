/* *
 * Counting Sundays
 * ----------------------
 * You are given the following information, but you may prefer to do some research for yourself.
 * 1 Jan 1900 was a Monday.
 * Thirty days has September,
 * April, June and November.
 * All the rest have thirty-one,
 * Saving February alone,
 * Which has twenty-eight, rain or shine.
 * And on leap years, twenty-nine.
 * A leap year occurs on any year evenly divisible by 4, but not on a century
 * unless it is divisible by 400.
 *
 * How many Sundays fell on the first of the month during the twentieth
 * century (1 Jan 1901 to 31 Dec 2000)?
 * ------------------------------------------------------------------
 */

fn is_leap_year(year: u32) -> bool {
    (year % 4 == 0 && year % 100 != 0) || year % 400 == 0
}

fn days_in_year(year: u32) -> u32 {
    if is_leap_year(year) {366} else {365}
}

fn month_first_is_sunday_in_year(year: u32, first_day_of_the_year: u32) -> u32{
    assert!(first_day_of_the_year < 7);
    let mut result: u32 = 0;
    let mut day_in_year: u32 = first_day_of_the_year;
    let mut day_in_months: [u32; 12] = [31,28,31,30,31,30,31,31,30,31,30,31];
    if is_leap_year(year) {day_in_months[1]=29;}

    for day_in_month in day_in_months {
        if day_in_year % 7 == 5 {
            result += 1;
        }
        day_in_year += day_in_month;
    }

    result
}

fn main() {
    let mut result: u32 = 0;
    let mut first_day_of_the_year: u32 = 1;
    for year in 1901..2001{
        result += month_first_is_sunday_in_year(year, first_day_of_the_year);
        first_day_of_the_year += days_in_year(year);
        first_day_of_the_year = first_day_of_the_year % 7;
    }
    println!("Sundays fell on the first of the month during the twentieth centaury (1 Jan 1901 to 31 Dec 2000): {}", result);
}
