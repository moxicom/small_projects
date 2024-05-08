#include <iostream>
#include <vector>
#include <cmath>

using namespace std;

int main()
{
    vector<int> grades(26 + 1, 0);
    vector<char> dict(26 + 1, 'A');

    string s;
    cin >> s;
    int l = s.length();

    char el = 'Z';
    char bad = 'A';

    for (int i = 1; i <= 26; i++)
    {
        dict[i] = el;
        el--;
    }
    // for (auto now : dict)
    // {
    //     cout << now << ' ';
    // }
    // cout << '\n';

    for (int i = 0; i < l; i++)
    {
        char ch = s[i];
        if (ch > bad)
        {
            bad = ch;
        }
        grades['Z' - ch + 1] += 1;
    }

    int sum = 0;

    for (int i = 1; i < 27; i++)
    {
        sum += (grades[i]) * i;
    }

    // cout << sum << endl;

    int ost = round(float(sum) / l);
    int mod = sum % l;

    // cout << ost << endl;
    if (bad - dict[ost] > 1)
    {
        // cout << "asdasd" << endl;
        cout << char(bad - 1);
    }
    else
    {
        cout << dict[ost];
    }

    return 0;
}