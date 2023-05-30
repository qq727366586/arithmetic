package main

import (
	"fmt"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

// 前序遍历
func PreorderTraversal(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Data) // 输出当前节点的值
	PreorderTraversal(root.Left)
	PreorderTraversal(root.Right)
}

// 中序遍历
func InorderTraversal(root *Node) {
	if root == nil {
		return
	}

	InorderTraversal(root.Left)  // 递归遍历左子树
	fmt.Printf("%d ", root.Data) // 输出当前节点的值
	InorderTraversal(root.Right) // 递归遍历右子树
}

// 后序遍历
func PostorderTraversal(root *Node) {
	if root == nil {
		return
	}

	PostorderTraversal(root.Left)  // 递归遍历左子树
	PostorderTraversal(root.Right) // 递归遍历右子树
	fmt.Printf("%d ", root.Data)   // 输出当前节点的值
}

func main() {
	// 构建二叉树
	root := &Node{Data: 1}
	root.Left = &Node{Data: 2}
	root.Right = &Node{Data: 3}
	root.Left.Left = &Node{Data: 4}
	root.Left.Right = &Node{Data: 5}
	root.Right.Left = &Node{Data: 6}
	root.Right.Right = &Node{Data: 7}

	fmt.Println("前序遍历:")
	PreorderTraversal(root)
	fmt.Println()

	fmt.Println("中序遍历:")
	InorderTraversal(root)
	fmt.Println()

	fmt.Println("后序遍历:")
	PostorderTraversal(root)
	fmt.Println()
}
