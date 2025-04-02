/* *
 * Multiples of 3 or 5
 * -------------------
 * If we list all the natural numbers below 10
 * that are multiples of 3 or 5, we get 3, 5,
 * 6 and 9. The sum of these multiples is 23.
 *
 * Find the sum of all the multiples of 3 or 5 below 1000.
 * ---------------------------------------------------------
 */
#include <iostream>
#include <vector>

#define LOWER_LIMIT 1
#define UPPER_LIMIT 1000

int main(int argc, char **argv)
{
    int sum = 0;
    std::vector<int> result;
    for (int i = LOWER_LIMIT; i < UPPER_LIMIT; i++)
    {
        if (i % 3 == 0 || i % 5 == 0)
        {
            result.push_back(i);
            sum += i;
        }
    }

    int i = 0;
    for (std::vector<int>::iterator it = result.begin();
         it != result.end();
         it++)
    {
        std::cout << *it << " ";

        if (i % 16 == 15)
        {
            std::cout << "\n";
        }
        else if (i % 8 == 7)
        {
            std::cout << "   ";
        }

        i++;
    }
    std::cout << "\n"
              << "Vector size: " << result.size() << "\n"
              << "Vector sum:" << sum << "\n";

    return 0;
}
