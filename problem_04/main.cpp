/* *
 * Largest Palindrome Product
 * ----------------------
 * A palindromic number reads the same both ways. The largest palindrome
 * made from the product of two 2-digit numbers is 9009 = 91 * 99.
 *
 * Find the largest palindrome made from the product of two 3-digit numbers.
 * ------------------------------------------------------------------
 */
#include <iostream>
#include <cassert>
#include <vector>

#define NUMBER 9009

bool is_palindrom(unsigned int n)
{
    auto str_n = std::to_string(n);
    size_t length = str_n.length();

    for (size_t i = 0; i < length / 2; ++i)
    {
        if (str_n[i] != str_n[length - i - 1])
        {
            return false;
        }
    }
    return true;
}

int main(int argc, char **argv)
{
    assert(is_palindrom(NUMBER) == true);
    unsigned int largest_palindrom = 0, largest_palindrom_prod_a = 0, largest_palindrom_prod_b = 0;
    std::vector<unsigned int> palindroms;
    for (unsigned int a = 999; a > 99; a--)
    {
        for (unsigned int b = 999; b > 99; b--)
        {
            unsigned int mult = a * b;
            if (is_palindrom(mult))
            {
                std::cout << mult << " = " << b << " * " << a << "\n";
                if (mult > largest_palindrom)
                {
                    largest_palindrom = mult;
                    largest_palindrom_prod_a = a;
                    largest_palindrom_prod_b = b;
                }
            }
        }
    }
    std::cout << "Largest Palindrome Product\n";
    std::cout << largest_palindrom << " = " << largest_palindrom_prod_b << " * " << largest_palindrom_prod_a << "\n";
    return 0;
}
