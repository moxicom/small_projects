#include <iostream>
#include <vector>
#include <map>
#include <queue>

using namespace std;

struct Path
{
    int to;
    int time;
    long long cost = 0; // 0 == old way
    int num = -1;
};

bool hasPath(map<int, vector<Path>> &ways, int start, int end, int maxTime, long long maxCost)
{
    queue<int> q;
    vector<bool> visited(ways.size() + 1, false); // Assuming vertices are numbered from 1 to n
    q.push(start);
    visited[start] = true;

    while (!q.empty())
    {
        int current = q.front();
        q.pop();

        for (const auto &path : ways[current])
        {
            if (!visited[path.to] && path.time <= maxTime)
            {
                if (path.to == end)
                    return true;
                q.push(path.to);
                visited[path.to] = true;
            }
        }
    }
    return false;
}

int main()
{
    int n, m;
    cin >> n >> m;

    map<int, vector<Path>> ways;

    map<int, map<int, int>> time; // need to complete

    vector<int> cost(n, -1);

    for (int i = 1; i < m + 1; i++)
    {
        int u, v, t;
        cin >> u >> v >> t;
        ways[u] = {{v, t, 0, -1}};
        ways[v] = {{u, t, 0, -1}};
    }

    int k;
    cin >> k;
    for (int i = 1; i < k + 1; i++)
    {
        int u, v, t, c;
        cin >> u >> v >> t >> c;
        ways[u].push_back({v, t, c, i});
        ways[v].push_back({u, t, c, i});
    }

    int p;
    cin >> p;
    vector<int> ans;
    for (int i = 0; i < p; i++)
    {
        int a, b, t;
        cin >> a >> b >> t;
        time[a][b] = t;
        time[b][a] = t;
        if (!hasPath(ways, a, b, t))
        {
            // Check if we can use k's
            cout << "No way in " << a << ' ' << b << endl;
        }
    }
    if (ans.size() == 0)
    {
        cout << 0;
    }
    else
    {
        cout << ans.size();
        for (int i = 0; i < ans.size(); i++)
        {
            cout << ans[i] << " ";
        }
    }
}