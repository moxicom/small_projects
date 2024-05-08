package main

import (
	"fmt"
	"os"
	"sort"
)

type Info struct {
	idx    int
	maxVal int
}

func main() {
	var n, m int
	fmt.Scanf("%v %v", &n, &m)

	max := m
	v := make([]Info, n, n)
	ans := make([]int, n, n)

	for i := 0; i < n; i++ {
		var ai int
		fmt.Scan(&ai)
		v[i] = Info{i, ai}
	}

	sort.Slice(v, func(i, j int) bool {
		return v[i].maxVal > v[j].maxVal
	})

	for _, val := range v {
		if max-val.maxVal <= 0 {
			fmt.Println(-1)
			os.Exit(0)
		}
		ans[val.idx] = max
		max--
	}

	for _, v := range ans {
		fmt.Printf("%v ", v)
	}

}

