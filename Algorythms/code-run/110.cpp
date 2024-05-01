// 110 - scanline

#include <iostream>
#include <vector>
#include <map>
#include <algorithm>

struct event
{
    int x;
    int type;
};

void print_vector(std::vector<event> &v)
{
    std::cout << std::endl;
    for (int i = 0; i < v.size(); i++)
    {
        std::cout << v[i].x << " " << v[i].type << " ";
        std::cout << std::endl;
    }
    std::cout << std::endl;
}

int main()
{
    int n;
    int m;
    std::cin >> n >> m;
    std::vector<event> points;

    for (int i = 0; i < m; i++)
    {
        int a, b;
        std::cin >> a >> b;
        if (a > b)
        {
            std::swap(a, b);
        }
        points.push_back({a, 1});
        points.push_back({b, -1});
    }

    std::sort(points.begin(), points.end(), [](const event &e1, const event &e2)
              { return e1.x == e2.x ? e1.type > e2.type : e1.x < e2.x; });

    // print_vector(points);

    int uncovered_students = 0;
    int current_observers = 0;
    int prev_end = -1;
    for (auto now : points)
    {
        current_observers += now.type;
        if (current_observers == 1 && now.type == 1)
        {
            uncovered_students += now.x - prev_end - 1;
            // std::cout << now.x << " " << uncovered_students << std::endl;
        }
        if (now.type == -1)
        {
            prev_end = now.x;
        }
    }

    uncovered_students += n - prev_end - 1;

    std::cout << uncovered_students << std::endl;

    return 0;
}