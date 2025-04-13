/* *
 * Maximum Path Sum I
 * ----------------------
 * By starting at the top of the triangle below and moving to
 * adjacent numbers on the row below, the maximum total from
 * top to bottom is 23.
 *                               '3'
 *                             '7'  4
 *                            2  '4'  6
 *                           8   5 '9'  3
 *
 * That is 3 + 7 + 4 + 9 = 23
 *
 * Find the maximum total from top to bottom of the triangle below:
 * ------------------------------------------------------------------
 */

#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <sstream>
#include <assert.h>
#define DATA_PATH_TEST "./problem_067/data_test.txt"
#define DATA_PATH "./problem_067/data.txt"

void display_vector(std::vector<unsigned int> v, std::string name = "")
{
    if (name != "")
    {
        std::cout << "Vector: " << name << "\n";
    }
    for (size_t i = 0; i < v.size(); ++i)
    {
        std::cout << v[i] << ' ';
    }
    std::cout << '\n';
}

std::vector<std::vector<unsigned int>> read_data(std::string path)
{
    std::ifstream file(path);
    if (!file.is_open())
    {
        return std::vector<std::vector<unsigned int>>{};
    }
    std::string line;
    std::vector<std::vector<unsigned int>> all_numbers;

    while (std::getline(file, line))
    {
        std::vector<unsigned int> line_numbers;
        std::istringstream iss(line);
        int number;

        while (iss >> number)
        {
            line_numbers.push_back(number);
        }
        all_numbers.push_back(line_numbers);
    }

    file.close();
    return all_numbers;
}

unsigned int max_path_sum(std::vector<std::vector<unsigned int>> data)
{
    size_t n = data.size();
    assert(n > 1);

    std::vector<unsigned int> curr_sum(data[n - 1]);
    // std::vector<unsigned int> temp(curr_sum);
    display_vector(curr_sum, "curr_sum");

    for (size_t i = n - 2;; i--)
    {
        display_vector(data[i], ("data[]"));
        std::vector<unsigned int> temp(data[i]);
        for (size_t j = 0; j <= i; ++j)
        {
            temp[j] += curr_sum[j] > curr_sum[j + 1] ? curr_sum[j] : curr_sum[j + 1];
        }
        curr_sum = temp;
        if (i == 0)
        {
            break;
        }
    }
    display_vector(curr_sum, "curr_sum");
    return curr_sum[0];
}

int main()
{
    auto test_data = read_data(DATA_PATH_TEST);
    auto test_result = max_path_sum(test_data);
    assert(test_result == 23);

    auto data = read_data(DATA_PATH);
    auto result = max_path_sum(data);
    std::cout << "Maximum total from top to bottom of the triangle is: " << result << "\n";
    return 0;
}
