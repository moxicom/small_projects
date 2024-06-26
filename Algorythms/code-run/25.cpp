//
// 25
//
#include <iostream>
#include <vector>
#include <map>
#include <algorithm>

void print_vector(std::vector<int> &v)
{
    for (int i = 0; i < v.size(); i++)
    {
        std::cout << v[i] << " ";
    }
    std::cout << std::endl;
}

int main()
{
    int n;
    std::cin >> n;
    std::vector<std::vector<int>> v;

    for (int i = 0; i < n; i++)
    {
        int num;
        std::cin >> num;
        int maxLen = 0;
        int idx = 0;
        int v_size = v.size();
        for (int k = 0; k < v_size; k++)
        {
            int vk_len = v[k].size();
            int tempLen = 0;
            for (int j = 0; j < vk_len; j++)
            {
                if (v[k][j] >= num)
                {
                    break;
                }
                tempLen = j + 1;
            }
            if (tempLen > maxLen)
            {
                maxLen = tempLen;
                idx = k;
            }
        }

        // std::cout << "maxlen: " << maxLen << std::endl;
        if (maxLen == 0)
        {
            v.push_back({num});
            // std::cout << "Created new, " << num << std::endl;
        }
        else
        {
            if (maxLen == v[idx].size())
            {
                v[idx].push_back(num);

                // std::cout << "Added to existing, " << idx << " " << num << std::endl;
                // print_vector(v[idx]);
            }
            else
            {
                std::vector<int> temp_new(maxLen);
                std::copy(&v[idx][0], &v[idx][maxLen], temp_new.begin());
                temp_new.push_back(num);

                v.push_back(temp_new);

                // std::cout << "Copied a slice, " << idx << " " << num << std::endl;
                // print_vector(v[v_size]);
            }
        }
    }

    int maxLenIdx = -1;
    int maxLen = 0;

    for (int i = 0; i < v.size(); i++)
    {
        int len = v[i].size();
        if (len > maxLen)
        {
            maxLen = len;
            maxLenIdx = i;
        }
    }

    for (int i = 0; i < v[maxLenIdx].size(); i++)
    {
        std::cout << v[maxLenIdx][i] << " ";
    }
    return 0;
}