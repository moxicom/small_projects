package algs

import (
	"container/list"
	"fmt"
)

// node is a graph node
type node struct {
	Id      string
	friends map[string]*node
}

// Nodes returns a list of all graph nodes
func Nodes(n *node) []*node {
	//track the visited nodes
	visited := make(map[string]*node)
	// queue of the nodes to visit
	queue := list.New()
	queue.PushBack(n)
	// add the root node to the map of the visited nodes
	visited[n.Id] = n

	for queue.Len() > 0 {
		qnode := queue.Front()
		// iterate through all of its friends
		if qnode.Value.(*node).Id == "node5" {
			fmt.Println("Found it")
			break
		}
		// mark the visited nodes; enqueue the non-visted
		for id, node := range qnode.Value.(*node).friends {
			if _, ok := visited[id]; !ok {
				visited[id] = node
				queue.PushBack(node)
			}
		}
		queue.Remove(qnode)
	}

	nodes := make([]*node, 0)
	// collecte all the nodes into slice
	for _, node := range visited {
		nodes = append(nodes, node)
	}

	return nodes
}

func CheckAlg() {
	node1 := &node{Id: "node1", friends: make(map[string]*node)}
	node2 := &node{Id: "node2", friends: make(map[string]*node)}
	node3 := &node{Id: "node3", friends: make(map[string]*node)}

	// node2 is directly connected to node1
	node2.friends[node1.Id] = node1
	// node3 is directly connected to node2
	node3.friends[node2.Id] = node2
	// node1 is directly connected to node3
	node1.friends[node3.Id] = node3

	// root node; this graph is actually not a tree  so it does not have a root node
	root := &node{Id: "root", friends: make(map[string]*node)}
	n := make(map[string]*node)
	n[root.Id] = root
	n[root.Id].friends[node1.Id] = node1
	n[root.Id].friends[node2.Id] = node2
	n[root.Id].friends[node3.Id] = node3

	nodes := Nodes(root)
	for _, node := range nodes {
		fmt.Printf("%+v\n", node.Id)
	}
}
