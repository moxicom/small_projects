/*
Игрушечный лабиринт представляет собой прозрачную плоскую прямоугольную коробку, внутри
которой есть препятствия и перемещается шарик. Лабиринт можно наклонять влево, вправо, к
себе или от себя, после каждого наклона шарик перемещается в заданном направлении до
ближайшего препятствия или до стенки лабиринта, после чего останавливается. Целью игры
является загнать шарик в одно из специальных отверстий – выходов. Шарик проваливается в
отверстие, если оно встречается на его пути (шарик не обязан останавливаться в отверстии).

Первоначально шарик находится в левом верхнем углу лабиринта. Гарантируется, что решение
существует и левый верхний угол не занят препятствием или отверстием.

Формат ввода
В первой строке входного файла записаны числа N и M – размеры лабиринта (целые положительные
числа, не превышающие 100). Затем идет N строк по M чисел в каждой – описание лабиринта.
Число 0 в описании означает свободное место, число 1 – препятствие, число 2 – отверстие.

Формат вывода
Выведите единственное число – минимальное количество наклонов, которые необходимо сделать,
чтобы шарик покинул лабиринт через одно из отверстий.

input
4 5
0 0 0 0 1
0 1 1 0 2
0 2 1 0 0
0 0 1 0 0

output
3

*/

#include <iostream>
#include <vector>
#include <queue>
#include <algorithm>

using namespace std;

int n, m;

struct Event
{
    int moves;
    int x;
    int y;
};

pair<int, int> makeMoves(Event top, int x_incr, int y_incr, vector<vector<int>> v)
{
    int new_x = top.x;
    int new_y = top.y;
    for (;;)
    {
        new_x += x_incr;
        new_y += y_incr;
        if (new_x < 0 || new_y < 0 || new_x == n || new_y == m)
        {
            new_x -= x_incr;
            new_y -= y_incr;
            break;
        }
        else if (v[new_x][new_y] == 1)
        {
            new_x -= x_incr;
            new_y -= y_incr;
            break;
        }
        else if (v[new_x][new_y] == 2)
        {
            break;
        }
    }
    return {new_x, new_y};
}

void addMove(pair<int, int> res, int moves, queue<Event> &q, vector<pair<int, int>> &visited)
{
    if (find(visited.begin(), visited.end(), res) == visited.end())
    {
        visited.push_back(res);
        q.push({moves, res.first, res.second});
    }
}

int main()
{
    cin >> n >> m;
    vector<vector<int>> v(n);

    for (int i = 0; i < n; i++)
    {
        vector<int> row(m);
        for (int j = 0; j < m; j++)
        {
            // 0 - empty, 1 - wall, 2 - exit
            cin >> row[j];
        }
        v[i] = row;
    }

    queue<Event> q;
    vector<pair<int, int>> visited;
    q.push({0, 0, 0});

    for (; !q.empty();)
    {
        Event top = q.front();
        q.pop();
        visited.push_back({top.x, top.y});

        // if found an exit
        if (v[top.x][top.y] == 2)
        {
            cout << top.moves;
            return 0;
        }

        pair<int, int> res;
        // go bottom (x+1)
        res = makeMoves(top, 1, 0, v);
        addMove(res, top.moves + 1, q, visited);
        // go top (x-1)
        res = makeMoves(top, -1, 0, v);
        addMove(res, top.moves + 1, q, visited);
        // go right (y+1)
        res = makeMoves(top, 0, 1, v);
        addMove(res, top.moves + 1, q, visited);
        // go left (y-1)
        res = makeMoves(top, 0, -1, v);
        addMove(res, top.moves + 1, q, visited);
    }

    cout << -1;

    return 0;
}
