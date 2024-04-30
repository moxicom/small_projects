/*
Между некоторыми деревнями края Васюки ходят автобусы. Поскольку пассажиропотоки здесь не очень большие,
то автобусы ходят всего несколько раз в день.

Марии Ивановне требуется добраться из деревни d в деревню v как можно быстрее (считается, что в момент
времени 0 она находится в деревне d).

Формат ввода
Сначала вводится число N – общее число деревень (1 ≤ N ≤ 100), затем номера деревень d и v, за ними следует
количество автобусных рейсов R (0 ≤ R ≤ 10000). Далее идут описания автобусных рейсов. Каждый рейс задается
номером деревни отправления, временем отправления, деревней назначения и временем прибытия (все времена –
целые от 0 до 10000). Если в момент t пассажир приезжает в какую-то деревню, то уехать из нее он может в
любой момент времени, начиная с t.

Формат вывода
Выведите минимальное время, когда Мария Ивановна может оказаться в деревне v. Если она не сможет с помощью
указанных автобусных рейсов добраться из d в v, выведите -1.

Выходные данные
3
1 3
4
1 0 2 5
1 1 2 3
2 3 3 5
1 1 3 10

Входные данные
5

*/

#include <iostream>
#include <map>
#include <limits>
#include <vector>
#include <queue>

struct Path
{
    int to;
    int startTime;
    int endTime;
};

long dijkstra_algorithm(std::map<int, std::vector<Path>> &m, int nodesAmount, int start, int finish)
{
    const long INF = std::numeric_limits<int>::max();
    std::vector<bool> visited(nodesAmount + 1, false);
    std::vector<long> dist(nodesAmount + 1, INF);

    long minIdx = start;
    dist[minIdx] = 0;

    for (;;)
    {
        visited[minIdx] = true;
        for (auto path : m[minIdx])
        {
            if (path.startTime >= dist[minIdx])
            {
                // update a time
                dist[path.to] = dist[path.to] < path.endTime ? dist[path.to] : path.endTime;
            }
        }

        int minvalue = INF;

        for (int i = 0; i < nodesAmount + 1; i++)
        {
            if (!visited[i] && dist[i] < minvalue)
            {
                minvalue = dist[i];
                minIdx = i;
            }
        }

        if (minvalue == INF)
            break;
    }

    if (dist[finish] != INF)
    {
        return dist[finish];
    }
    else
    {
        return -1;
    }
}

int main()
{
    int start, finish, nodesAmount, pathsAmount;

    std::cin >> nodesAmount;
    std::cin >> start >> finish;
    std::cin >> pathsAmount;

    std::map<int, std::vector<Path>> m;

    for (long i = 0; i < pathsAmount; i++)
    {
        // FROM_P FROM_T TO_P TO_T
        int f_p, f_t, t_p, t_t;
        std::cin >> f_p >> f_t >> t_p >> t_t;
        m[f_p].push_back({t_p, f_t, t_t});
    }

    std::cout << dijkstra_algorithm(m, nodesAmount, start, finish);
    return 0;
}