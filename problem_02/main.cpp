/* *
 * Even Fibonacci Numbers
 * ----------------------
 * Each new term in the Fibonacci sequence is generated
 * by adding the previous two terms. By starting with 1
 * and 2, the first 10 terms will be:
 * 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
 *
 * By considering the terms in the Fibonacci sequence
 * whose values do not exceed four million, find the
 * sum of the even-valued terms.
 * ------------------------------------------------------------------
 */
#include <iostream>
#include <vector>

#define F_0 1
#define F_1 2
#define LIMIT 4000000

int main(int argc, char **argv)
{
    std::vector<unsigned long long> fibbonaci = {F_0, F_1};
    std::cout << fibbonaci[0] << "\n";
    std::cout << fibbonaci[1] << "\n";
    while (true)
    {
        size_t n = fibbonaci.size();
        unsigned long long next_val = fibbonaci[n - 2] + fibbonaci[n - 1];
        std::cout << next_val << "\n";
        fibbonaci.push_back(next_val);
        if (next_val > LIMIT)
        {
            break;
        }
    }

    unsigned long long even_sum = 0, odd_sum = 0;
    for (auto it = fibbonaci.begin(); it != fibbonaci.end(); ++it)
    {
        if (*it % 2 == 0)
        {
            even_sum += *it;
        }
        else
        {
            odd_sum += *it;
        }
    }
    std::cout << "Even sum:" << even_sum << "\n";
    std::cout << "Odd sum:" << odd_sum << "\n";
    return 0;
}
