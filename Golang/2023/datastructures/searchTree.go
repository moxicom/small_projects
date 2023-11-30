package datastructures

import "fmt"

type searchTree struct {
	head *searchTreeNode
}

type searchTreeNode struct {
	// data. key and value are just examples
	key   int
	value int
	// childs
	left  *searchTreeNode
	right *searchTreeNode
}

func NewSearchTree(key, value int) *searchTree {
	return &searchTree{
		head: &searchTreeNode{
			key:   key,
			value: value,
		},
	}
}

func (t *searchTree) Insert(key, value int) {
	insertRec(t.head, key, value)
}

func insertRec(node *searchTreeNode, key, value int) {
	if key < node.key {
		if node.left == nil {
			node.left = &searchTreeNode{
				key:   key,
				value: value,
			}
		} else {
			insertRec(node.left, key, value)
		}
	} else if key >= node.key {
		if node.right == nil {
			node.right = &searchTreeNode{
				key:   key,
				value: value,
			}
		} else {
			insertRec(node.right, key, value)
		}
	}
}

// returns node
func (t *searchTree) Search(key int) *searchTreeNode {
	return searchRec(t.head, key)
}

func searchRec(node *searchTreeNode, key int) *searchTreeNode {
	if node == nil {
		return nil
	}
	if node.key == key {
		return node
	}
	if node.key < key {
		return searchRec(node.right, key)
	} else {
		return searchRec(node.left, key)
	}

}

// returns a node with the smallest key
func (t *searchTree) GetMin() *searchTreeNode {
	return getMinRec(t.head)
}

func getMinRec(node *searchTreeNode) *searchTreeNode {
	if node.left == nil {
		return node
	}
	return getMinRec(node.left)
}

// returns a node with the biggest key
func (t *searchTree) GetMax() *searchTreeNode {
	return getMaxRec(t.head)
}

func getMaxRec(node *searchTreeNode) *searchTreeNode {
	if node.right == nil {
		return node
	}
	return getMaxRec(node.right)
}

func (t *searchTree) Delete(key int) *searchTreeNode {
	return deleteRec(t.head, key)
}

func deleteRec(node *searchTreeNode, key int) *searchTreeNode {
	if node == nil {
		return nil
	} else if node.key > key {
		node.left = deleteRec(node.left, key)
	} else if node.key < key {
		node.right = deleteRec(node.right, key)
	} else {
		if node.left == nil && node.right != nil {
			node = node.right
		} else if node.left != nil && node.right == nil {
			node = node.left
		} else if node.left == nil && node.right == nil {
			node = nil
		} else {
			// we could get max left or min right
			maxLeft := getMaxRec(node.left)
			node.key = maxLeft.key
			node.value = maxLeft.value
			node.left = deleteRec(node.left, node.key)
		}
	}
	return node
}

func (t *searchTree) PrintInOrder() {
	if t == nil {
		return
	}
	printInOrderRec(t.head)
}

func printInOrderRec(node *searchTreeNode) {
	if node.left != nil {
		printInOrderRec(node.left)
	}
	fmt.Printf("key: %v, value: %v\n", node.key, node.value)
	if node.right != nil {
		printInOrderRec(node.right)
	}
}

func (t *searchTree) DeleteTree() {
	deleteTree(t.head)
}

func deleteTree(node *searchTreeNode) {
	if node == nil {
		return
	}
	deleteTree(node.left)
	deleteTree(node.right)
}
