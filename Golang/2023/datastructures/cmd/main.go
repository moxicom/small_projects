package main

import (
	"datastructures"
)

func main() {
	tree := datastructures.NewSearchTree(0, 121)
	tree.Insert(-21, 123123)
	tree.Insert(212, 1223)
	tree.Insert(11, 1333)
	tree.Delete(0)
	tree.PrintInOrder()

}
