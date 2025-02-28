// Question 2
package main

import "fmt"

type ExprNode struct {
    Value string
    Left  *ExprNode
    Right *ExprNode
}

func Preorder(node *ExprNode) {
    if node == nil {
        return
    }
    fmt.Print(node.Value, " ")
    Preorder(node.Left)
    Preorder(node.Right)
}

func Postorder(node *ExprNode) {
    if node == nil {
        return
    }
    Postorder(node.Left)
    Postorder(node.Right)
    fmt.Print(node.Value, " ")
}

func main() {
    root := &ExprNode{Value: "+", Left: &ExprNode{Value: "a"}, Right: &ExprNode{Value: "-"}}
    root.Right.Left = &ExprNode{Value: "b"}
    root.Right.Right = &ExprNode{Value: "c"}

    fmt.Print("Preorder Traversal: ")
    Preorder(root)
    fmt.Println()

    fmt.Print("Postorder Traversal: ")
    Postorder(root)
    fmt.Println()
}