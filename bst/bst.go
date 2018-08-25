package main

import (
	"fmt"
)

// Node -- Represents the basic node in tree
type Node struct {
	data  int
	left  *Node
	right *Node
}

// IterativeInsert -- Insert the value into the BST
func IterativeInsert(root *Node, v int) *Node {
	if root == nil {
		root = &Node{data: v, right: nil, left: nil}
	} else if v >= root.data {
		root.right = IterativeInsert(root.right, v)
	} else {
		root.left = IterativeInsert(root.left, v)
	}
	return root
}

// InOrder -- Print the BST in-order
func InOrder(root *Node) {
	if root == nil {
		return
	}
	InOrder(root.left)
	fmt.Printf("%d ", root.data)
	InOrder(root.right)
}

// PreOrder -- Print the BST in pre-order
func PreOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.data)
	PreOrder(root.left)
	PreOrder(root.right)
}

// PostOrder -- Print the BST in post order
func PostOrder(root *Node) {
	if root == nil {
		return
	}
	PostOrder(root.left)
	PostOrder(root.right)
	fmt.Printf("%d ", root.data)
}

// GetMin -- Get the minimum value of the BST
func GetMin(root *Node) *Node {
	if root == nil || root.left == nil {
		return root
	}
	return GetMin(root.left)
}

// GetMax -- Get the maximum value of the BST
func GetMax(root *Node) *Node {
	if root == nil || root.right == nil {
		return root
	}
	return GetMax(root.right)
}

func main() {
	var root *Node
	nums := []int{
		1, 34, 52, 6, 23, 62, 7, 78, 2,
	}
	for _, each := range nums {
		root = IterativeInsert(root, each)
	}
	InOrder(root)
	fmt.Println()
	PreOrder(root)
	fmt.Println()
	PostOrder(root)
	fmt.Println()
	fmt.Println(GetMin(root).data)
	fmt.Println(GetMax(root).data)
}
