/*
112
https://coderun.yandex.ru/problem/seating-arrangements/description?currentPage=1&difficulty=MEDIUM&groups=%D0%90%D0%BB%D0%B3%D0%BE%D1%80%D0%B8%D1%82%D0%BC%D1%8B&pageSize=10&rowNumber=36

112. Рассадка в аудитории
Экзамен по берляндскому языку проходит в узкой и длинной аудитории. На экзамен пришло N студентов.
Все они посажены в ряд. Таким образом, позиция каждого человека задается координатой на оси Ox
(эта ось ведёт вдоль длинной аудитории). Два человека могут разговаривать, если расстояние между
ними меньше или равно D. Какое наименьшее количество типов билетов должен подготовить преподаватель,
чтобы никакие два студента с одинаковыми билетами не могли разговаривать? Выведите способ раздачи
преподавателем билетов.

Формат ввода
В первой строке входного файла содержится два целых числа N, D (1 ≤  N ≤  10000; 0 ≤ D ≤ 10^6).
Вторая строка содержит последовательность различных целых чисел 𝑋1  ... , обозначает координату
вдоль оси Ox i-го студента

Формат вывода
В первую строчку выходного файла выведите количество вариантов, а во вторую, разделяя пробелами,
номера вариантов студентов в том порядке, в каком они перечислены во входном файле.



*/

#include <iostream>
#include <algorithm>
#include <vector>
#include <map>

using namespace std;

bool compareSecond(const pair<long, long> a, const pair<long, long> b)
{
    return a.second < b.second;
}

int main()
{
    long n, d;
    cin >> n >> d;
    vector<pair<long, long>> v(n);
    vector<long> ends;
    map<long, long> ans;

    for (long i = 0; i < n; i++)
    {
        long val;
        cin >> val;
        v[i] = {val, i};
    }

    sort(v.begin(), v.end()); // sort by value inc

    ends.push_back(v[0].first + d);
    ans[v[0].second] = 1;

    for (long i = 1; i < n; i++)
    {
        bool hasLower = false;
        long currentVal = v[i].first;
        long currentIdx = v[i].second;
        for (int j = 0; j < ends.size(); j++)
        {
            if (ends[j] < currentVal)
            {
                hasLower = true;
                ends[j] = currentVal + d;
                ans[currentIdx] = j + 1;
                break;
            }
        }

        if (!hasLower)
        {
            ends.push_back(currentVal + d);
            ans[currentIdx] = ends.size();
        }
    }

    sort(v.begin(), v.end(), compareSecond);
    cout << ends.size() << endl;
    for (int i = 0; i < n; i++)
    {
        cout << ans[v[i].second] << " ";
    }

    return 0;
}