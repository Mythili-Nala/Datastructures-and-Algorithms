package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func (n *Node) Init(data int) *Node {
	n.data = data
	n.left = nil
	n.right = nil
	return n
}

func New(data int) *Node {
	return new(Node).Init(data)
}

func DeleteTree(node *Node) {
	if node == nil {
		return
	}
	DeleteTree(node.left)
	DeleteTree(node.right)
	fmt.Printf("\nDeleting node: %d", node.data)
	node = nil
}

func main() {

	root := New(1)

	root.left = New(2)
	root.right = New(3)

	root.left.left = New(4)
	root.left.right = New(5)

	DeleteTree(root)
	root = nil

	fmt.Println("\nTree deleted")
}