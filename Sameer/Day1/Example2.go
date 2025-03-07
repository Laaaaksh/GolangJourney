package main

import "fmt"

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

func Preorder(node *Node) string {
	if node == nil {
		return ""
	}
	return node.Value + " " + Preorder(node.Left) + Preorder(node.Right)
}

func Postorder(node *Node) string {
	if node == nil {
		return ""
	}
	return Postorder(node.Left) + Postorder(node.Right) + node.Value + " "
}

func main() {
	root := &Node{Value: "+", Left: &Node{Value: "a"}, Right: &Node{Value: "-", Left: &Node{Value: "b"}, Right: &Node{Value: "c"}}}
	fmt.Println("Preorder:", Preorder(root))
	fmt.Println("Postorder:", Postorder(root))
}
