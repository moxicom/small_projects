#include <iostream>
#include <vector>
#include <map>
// #include <algorithm>
#include <stack>

void dfs(int chosen, std::stack<int> &stack, std::map<int, std::vector<int>> &edges, std::map<int, char> &colors)
{
    colors[chosen] = 'G';
    for (auto now : edges[chosen])
    {
        if (colors[now] == 'W')
        {
            dfs(now, stack, edges, colors);
        }
        if (colors[now] == 'G')
        {
            std::cout << "-1";
            exit(0);
        }
    }
    colors[chosen] = 'B';
    stack.push(chosen);
}

int main()
{
    int N, M;
    std::cin >> N >> M;

    std::map<int, std::vector<int>> edges;
    std::map<int, char> colors;
    std::map<int, bool> end_v;
    std::stack<int> stack;

    for (int i = 1; i <= N; i++)
    {
        colors[i] = 'W';
        end_v[i] = false;
    }

    for (int i = 0; i < M; i++)
    {
        int nodeStart, nodeEnd;
        std::cin >> nodeStart >> nodeEnd;
        end_v[nodeEnd] = true;
        edges[nodeStart].push_back(nodeEnd);
    }

    for (int i = 1; i <= N; i++)
    {
        if (colors[i] == 'W' && !end_v[i])
        {
            // std::cout << "i: " << i << std::endl;
            dfs(i, stack, edges, colors);
        }
    }

    for (int i = 1; i <= N; i++)
    {
        if (colors[i] == 'W')
        {
            std::cout << "-1";
            return 0;
        }
    }

    for (; !stack.empty();)
    {
        std::cout << stack.top() << " ";
        stack.pop();
    }

    return 0;
}