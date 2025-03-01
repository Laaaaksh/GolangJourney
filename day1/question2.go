package main

import (
	"fmt"
)

// Node struct represents a node in the expression tree
type Node struct {
	Value string
	Left  *Node
	Right *Node
}

// Preorder traversal (Root -> Left -> Right)
func Preorder(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.Value, " ")
	Preorder(node.Left)
	Preorder(node.Right)
}

// Postorder traversal (Left -> Right -> Root)
func Postorder(node *Node) {
	if node == nil {
		return
	}
	Postorder(node.Left)
	Postorder(node.Right)
	fmt.Print(node.Value, " ")
}

func main() {
	// Manually constructing the expression tree for "a + b - c"
	root := &Node{Value: "+"}
	root.Left = &Node{Value: "a"}
	root.Right = &Node{Value: "-"}
	root.Right.Left = &Node{Value: "b"}
	root.Right.Right = &Node{Value: "c"}

	fmt.Println("Preorder Traversal:")
	Preorder(root)
	fmt.Println()

	fmt.Println("Postorder Traversal:")
	Postorder(root)
	fmt.Println()
}
