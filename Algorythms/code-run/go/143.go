package main

import "fmt"

type stack []int

func (s stack) pop() (stack, int) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) top() int {
	l := len(s)
	return s[l-1]
}

func (s stack) push(val int) stack {
	return append(s, val)
}

func main() {
	var n int
	fmt.Scan(&n)

	s := stack{}
	t := stack{}
	expected := 1

	var val int
	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		s = s.push(val)

		for len(s) > 0 && s.top() == expected {
			var topVal int
			s, topVal = s.pop()
			t = t.push(topVal)
			expected++
		}
	}

	if expected == n+1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
