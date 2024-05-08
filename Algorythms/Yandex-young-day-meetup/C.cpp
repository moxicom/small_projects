#include <iostream>
#include <string>
#include <unordered_map>

using namespace std;

bool canBeginGame(int k, const string &s, const string &t)
{
    // Проверяем, достаточно ли символов в строках для начала игры
    if (s.length() < k || t.length() < k)
    {
        return false;
    }

    // Для Саши: создаем отображение символов в их количестве в строке t
    unordered_map<char, int> charCount;
    for (char c : t)
    {
        charCount[c]++;
    }

    // Создаем sliding window для Феди
    unordered_map<char, int> window;
    for (int i = 0; i < k; ++i)
    {
        window[s[i]]++;
    }

    // Проверяем, достаточно ли символов в sliding window для Феди
    for (int i = k; i <= s.length(); ++i)
    {
        bool validWindow = true;
        for (const auto &[c, count] : window)
        {
            if (charCount[c] < count)
            {
                validWindow = false;
                break;
            }
        }
        if (validWindow)
        {
            return true;
        }
        if (i < s.length())
        {
            window[s[i]]++;
            window[s[i - k]]--;
        }
    }

    return false;
}

int main()
{
    int k;
    string s, t;

    // Ввод данных
    cin >> k >> s >> t;

    // Проверка возможности начала игры
    if (canBeginGame(k, s, t))
    {
        cout << "YES" << endl;
    }
    else
    {
        cout << "NO" << endl;
    }

    return 0;
}
