
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

func Search(root *Node, key int) *Node {
	if root == nil || root.data == key {
		//fmt.Println("The given previous node cannot be NULL")
		return root
	}

	if root.data < key {
		return Search(root.right, key)
	}
	return Search(root.left, key)
}

func PrintInOrder(root *Node) {
	if root != nil {
		PrintInOrder(root.left)
		fmt.Printf("%d \n", root.data)
		PrintInOrder(root.right)
	}
}

func Insert(node *Node, key int) *Node {
	if node == nil {
		return New(key)
	}

	if key < node.data {
		node.left = Insert(node.left, key)
	} else if key > node.data {
		node.right = Insert(node.right, key)
	}
	return node
}

func main() {

	root := New(50)
	Insert(root, 30)
	Insert(root, 20)
	Insert(root, 40)
	Insert(root, 70)
	Insert(root, 60)
	Insert(root, 80)
	PrintInOrder(root)
}