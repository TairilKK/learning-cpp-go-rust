/* *
 * Largest Prime Factor
 * ----------------------
 * The prime factors of 13195 are 5, 7, 13 and 29.
 *
 * What is the largest prime factor of the number 600851475143?
 * ------------------------------------------------------------------
 */
#include <iostream>
#include <cmath>
#include <vector>

#define NUMBER 600851475143

unsigned int is_prime(unsigned int n)
{
    unsigned int n_root = sqrt(n);

    for (unsigned int i = 2; i <= n_root; ++i)
    {
        if (n % i == 0)
        {
            return 0;
        }
    }
    return n;
}

std::vector<unsigned long> largest_prime_factor(unsigned long n)
{
    std::vector<unsigned long> vector;
    unsigned int n_root = sqrt(n);
    for (unsigned int i = 2; i <= n_root; ++i)
    {
        if (n % i == 0 && is_prime(i))
        {
            vector.push_back(i);
        }
    }

    return vector;
}

int main(int argc, char **argv)
{
    auto result = largest_prime_factor(NUMBER);
    for (auto it = result.begin(); it != result.end(); it++)
    {
        std::cout << *it << " ";
    }
    std::cout << "\n";
    return 0;
}
