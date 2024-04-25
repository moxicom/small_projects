/*
A. Дейкстра
Ограничение времени	1 секунда
Ограничение памяти	64Mb
Ввод	стандартный ввод или input.txt
Вывод	стандартный вывод или output.txt
Дан ориентированный взвешенный граф. Найдите кратчайшее расстояние от одной заданной вершины до другой.

Формат ввода
В первой строке содержатся три числа: N, S и F (1≤ N ≤ 100, 1 ≤ S, F ≤ N), где N — количество вершин графа, 
S — начальная вершина, а F — конечная.В следующих N строках вводится по N чисел, не превосходящих 100,
– матрица смежности графа, где -1 означает что ребра между вершинами нет, а любое неотрицательное число —
наличие ребра данного веса. На главной диагонали матрицы записаны нули.

Формат вывода
Выведите искомое расстояние или -1, если пути между указанными вершинами не существует.

Пример
Ввод	
3 2 1
0 1 1
4 0 1
2 1 0

Вывод
3
*/


#include <iostream>
#include <map>
#include <limits>
#include <vector>

int dijkstra_algorithm(std::map<int, std::vector<int>> &m, int nodesAmount, int start, int finish) {
    const int maxValue = std::numeric_limits<int>::max();
    std::vector<bool> visited(nodesAmount + 1, false);
    std::vector<int> dist(nodesAmount + 1, maxValue);

    int minIdx = start;
    dist[minIdx] = 0;

    for (;;)
    {
        visited[minIdx] = true;
        for (int i = 1; i < nodesAmount + 1; i++)
        {
            if (m[minIdx][i] != -1 && dist[i] > dist[minIdx] + m[minIdx][i])
                dist[i] = dist[minIdx] + m[minIdx][i];
        }

        int minValue = maxValue;

        for (int i = 1; i < nodesAmount + 1; i++)
        {
            if (dist[i] < minValue && !visited[i]){
                minIdx = i;
                minValue = dist[i];
            }
        }
        if (minValue == maxValue)
            break;
    }

    if (dist[finish] != maxValue)
        return dist[finish];
    else
        return -1;
}

int main()
{
    int n, s, f;
    std::cin >> n >> s >> f;
    std::map<int, std::vector<int>> m;

    for (int i = 1; i < n + 1; i++)
    {
        std::vector<int> relations(n + 1, -1);
        for (int j = 1; j < n + 1; j++)
        {
            int temp;
            std::cin >> temp;
            relations[j] = temp;
        }
        m[i] = relations;
    }

   std::cout << dijkstra_algorithm(m, n, s, f);
return 0;
}
