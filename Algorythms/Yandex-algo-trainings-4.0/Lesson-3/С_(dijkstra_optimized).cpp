/*
C. Быстрый алгоритм Дейкстры
Ограничение времени	5 секунд
Ограничение памяти	512Mb
Ввод	стандартный ввод или input.txt
Вывод	стандартный вывод или output.txt
Вам дано описание дорожной сети страны. Ваша задача – найти длину кратчайшего пути между городами А и B.

Формат ввода
Сеть дорог задана во входном файле следующим образом: первая строка содержит числа N и K (1 ≤ N ≤ 100000,
0 ≤ K ≤ 300000), где K – количество дорог. Каждая из следующих K строк содержит описание дороги с 
двусторонним движением – три целых числа ai, bi и li (1 ≤ ai, bi ≤ N, 1 ≤ li ≤ 106). Это означает, что
имеется дорога длины li, которая ведет из города ai в город bi. В последней строке находятся два числа А и
В – номера городов, между которыми надо посчитать кратчайшее расстояние (1 ≤ A, B ≤ N)

Формат вывода
Выведите одно число – расстояние между нужными городами. Если по дорогам от города А до города В доехать
 невозможно, выведите –1.
*/

#include <iostream>
#include <map>
#include <limits>
#include <vector>
#include <queue>

struct Path {
    long to;
    long weight;
};

struct Compare
{
    bool operator()(const std::pair<long, long> &a, const std::pair<long, long> &b) const
    {
        
        if (a.first == b.first)
            return a.second > b.second;
        return a.first > b.first;
    }
};

long dijkstra_algorithm(std::map<long, std::vector<Path>> &m, long nodesAmount, long start, long finish) {
    const long INF = std::numeric_limits<long>::max();
    std::vector<bool> visited(nodesAmount + 1, false);
    std::vector<long> dist(nodesAmount + 1, INF);
    std::priority_queue<std::pair<long, long>, std::vector<std::pair<long, long>>, Compare> pq;
    
    long minIdx = start;
    dist[minIdx] = 0;
    pq.push({0, start});

    while(!pq.empty()) {
        long u = pq.top().second; // vertex
        long d = pq.top().first; // distance
        pq.pop();
        visited[u] = true;

        if (visited[u]) {
            continue;
        }

        for (const auto &path : m[u]){
            if (dist[path.to] > d + path.weight){
                dist[path.to] = d + path.weight;
                pq.push({dist[path.to], path.to});
            }
        }        
    }

    if (dist[finish] != INF){
        return dist[finish];
    } else{
        return -1;
    }
}

int main()
{
    long n, k, start, finish;
    std::cin >> n >> k;
    std::map<long, std::vector<Path>> m;

    for (long i = 0; i < k; i++)
    {
        // FROM TO LENGTH
        long a, b, l;
        std::cin >> a >> b >> l;
        m[a].push_back({b, l});
        m[b].push_back({a, l});
    }

    std::cin >> start >> finish;

    std::cout << dijkstra_algorithm(m, n, start, finish);
    return 0;
}
