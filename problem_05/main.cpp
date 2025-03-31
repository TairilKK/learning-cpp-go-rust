/* *
 * Smallest Multiple
 * ----------------------
 * 2520 is the smallest number that can be divided
 * by each of the numbers from 1 to 10 without any remainder.
 *
 * What is the smallest positive number that is evenly
 * divisible by all of the numbers from 1 to 20?
 * ------------------------------------------------------------------
 */
#include <iostream>
#include <cmath>
#include <cassert>
#include <algorithm>
#include <vector>

void display_vector(std::vector<int> vec)
{
    for (auto it = vec.begin(); it != vec.end(); it++)
    {
        std::cout << *it << ' ';
    }
    std::cout << '\n';
}

std::vector<int>
number_product_rec(int n, std::vector<int> res)
{
    assert(n > 0);
    if (n == 1)
    {
        res.push_back(1);
        return res;
    }
    for (int i = 2; i < n; ++i)
    {
        if (n % i == 0)
        {
            res.push_back(i);
            return number_product_rec(n / i, res);
        }
    }
    res.push_back(1);
    res.push_back(n);
    return res;
}

std::vector<int>
number_product(int n)
{
    std::vector<int> result;
    result = number_product_rec(n, result);
    std::sort(result.begin(), result.end());
    return result;
}

std::vector<int>
concat_vectors(std::vector<int> vec, std::vector<int> sec)
{
    std::vector<int> temp = sec;
    for (auto it = vec.begin(); it != vec.end(); it++)
    {
        int i = 0;
        for (auto itt = temp.begin(); itt != temp.end(); ++itt, ++i)
        {
            if (*it == *itt)
            {
                temp.erase(temp.begin() + i);
                break;
            }
        }
    }

    for (auto it = temp.begin(); it != temp.end(); ++it)
    {
        vec.push_back(*it);
    }
    std::sort(vec.begin(), vec.end());
    return vec;
}

int main(int argc, char **argv)
{
    auto result = number_product(1);
    for (int i = 2; i < 21; i++)
    {
        auto current = number_product(i);
        result = concat_vectors(result, current);
    }
    int res = 1;
    for (auto it = result.begin(); it != result.end(); it++)
    {
        res *= *(it);
    }
    std::cout << "Result: " << res << '\n';
    return 0;
}
