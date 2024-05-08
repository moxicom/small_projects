#include <iostream>
#include <vector>
#include <cmath>
#include <queue>

using namespace std;

struct Info
{
    long wayTime;
    long speakTime;
};

int main()
{
    long n, T;
    cin >> n >> T;

    vector<Info> v(n);

    for (long i = 0; i < n; i++)
    {
        cin >> v[i].wayTime >> v[i].speakTime;
    }

    long maxPeopleAmount = 0;
    long peopleAmount = 0;
    priority_queue<long> pq;
    long allSum = 0;

    for (long i = 0; i < n; i++)
    {
        if (v[i].wayTime >= T)
        {
            break;
        }
        pq.push(v[i].speakTime);
        allSum += v[i].speakTime;
        peopleAmount++;

        long timeLeft = T - v[i].wayTime - allSum;

        if (timeLeft < 0)
        {
            // cout << '\n';
            while (!pq.empty())
            {
                timeLeft += pq.top();
                allSum -= pq.top();
                pq.pop();
                peopleAmount--;
                // cout << i << ' ' << allSum << ' ' << peopleAmount << '\n';
                if (timeLeft >= 0)
                {
                    break;
                }
            }
        }

        if (timeLeft >= 0 && peopleAmount > maxPeopleAmount)
        {
            maxPeopleAmount = peopleAmount;
        }
    }

    cout << maxPeopleAmount;

    return 0;
}
