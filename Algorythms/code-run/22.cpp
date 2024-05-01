// 22
// https://coderun.yandex.ru/problem/minimum-of-the-segment/description?currentPage=1&difficulty=MEDIUM&groups=%D0%90%D0%BB%D0%B3%D0%BE%D1%80%D0%B8%D1%82%D0%BC%D1%8B&pageSize=10&rowNumber=8
#include <iostream>
#include <deque>
#include <set>

int main()
{
    int n, k;
    std::multiset<int> s;
    std::cin >>
        n >> k;
    std::deque<int> d;
    for (int i = 0; i < k; i++)
    {
        int x;
        std::cin >> x;
        //
        s.insert(x);
        d.push_back(x);
    }

    std::cout << "" << *s.begin() << std::endl;

    for (int i = k; i < n; i++)
    {
        int x;
        std::cin >> x;
        //
        s.insert(x);
        d.push_back(x);
        s.erase(s.find(d.front()));
        d.pop_front();
        std::cout << "" << *s.begin() << std::endl;
    }

    return 0;
}