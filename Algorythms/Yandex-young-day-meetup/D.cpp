// timelimit 34

#include <iostream>
#include <vector>

using namespace std;

int main()
{
    int n, t;
    cin >> n >> t;

    vector<int> orders(n + 1);

    for (int i = 1; i <= n; i++)
    {
        cin >> orders[i];
    }

    for (int i = 0; i < t; i++)
    {
        int start, end;
        cin >> start >> end;

        int candidate = 0, count = 0;

        // Boyer-Moore Majority Vote Algorithm
        for (int j = start; j <= end; j++)
        {
            if (count == 0)
            {
                candidate = orders[j];
                count = 1;
            }
            else if (candidate == orders[j])
            {
                count++;
            }
            else
            {
                count--;
            }
        }

        // Count occurrences of the candidate
        count = 0;
        for (int j = start; j <= end; j++)
        {
            if (orders[j] == candidate)
            {
                count++;
            }
        }

        // Output the majority element or 0 if no majority
        if (count > (end - start + 1) / 2)
        {
            cout << candidate << endl;
        }
        else
        {
            cout << 0 << endl;
        }
    }

    return 0;
}
