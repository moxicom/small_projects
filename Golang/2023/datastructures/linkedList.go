package datastructures

// startTime := time.Now()

// myList := linkedList.NewLinkedList(12)

// myList.AddNum(22)
// myList.AddNum(23)

// fmt.Println(myList.GetAllNumbers()[:100])

// fmt.Println(time.Since(startTime))

type LinkedList struct {
	head *LinkedListNode
}

type LinkedListNode struct {
	Number int
	Next   *LinkedListNode
}

func NewLinkedList(num int) *LinkedList {
	return &LinkedList{
		head: &LinkedListNode{
			Number: num,
		}}
}

func (l *LinkedList) AddNum(num int) {
	l.SelectLastNode().Next = &LinkedListNode{
		Number: num,
	}
}

func (l *LinkedList) SelectLastNode() *LinkedListNode {
	node := l.head
	for node.Next != nil {
		node = node.Next
	}
	return node
}

func (l *LinkedList) GetAllNumbers() []int {
	node := l.head
	result := []int{}
	for node.Next != nil {
		result = append(result, node.Number)
		node = node.Next
	}
	result = append(result, node.Number)
	return result
}
