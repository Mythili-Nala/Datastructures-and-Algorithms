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

func PrintPostOrder(node *Node) {
	if node == nil {
		return
	}
	PrintPostOrder(node.left)
	PrintPostOrder(node.right)
	fmt.Printf("%d ", node.data)
}

func PrintInOrder(node *Node) {
	if node == nil {
		return
	}
	PrintInOrder(node.left)
	fmt.Printf("%d ", node.data)
	PrintInOrder(node.right)
}

func PrintPreOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.data)
	PrintPreOrder(node.left)
	PrintPreOrder(node.right)
}

func main() {
	root := New(1)

	root.left = New(2)
	root.right = New(3)

	root.left.left = New(4)
	root.left.right = New(5)

	fmt.Println("\nPreorder traversal of binary tree is :")
	PrintPreOrder(root)
	fmt.Println("\nInorder traversal of binary tree is :")
	PrintInOrder(root)
	fmt.Println("\nPostorder traversal of binary tree is :")
	PrintPostOrder(root)
}