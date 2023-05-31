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

// 前序遍历（非递归）
func preorderTraversal(root *Node) []int {
	if root == nil {
		return []int{}
	}
	stack, res := []*Node{root}, []int{}
	for len(stack) > 0 { // 栈中有节点时循环
		node := stack[len(stack)-1]  // 取出栈顶节点
		stack = stack[:len(stack)-1] // 栈顶出栈
		res = append(res, node.Data) // 将节点值加入结果数组
		if node.Right != nil {       // 右子节点不为空时入栈
			stack = append(stack, node.Right)
		}
		if node.Left != nil { // 左子节点不为空时入栈
			stack = append(stack, node.Left)
		}
	}
	return res
}

// 中序遍历（非递归）
func inorderTraversal(root *Node) []int {
	if root == nil {
		return []int{}
	}
	stack, res := []*Node{}, []int{}
	node := root
	for node != nil || len(stack) > 0 { // 当节点不为空或栈不为空时循环
		for node != nil { // 将左子节点一直入栈直到为空或者为叶子节点
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]   // 取出栈顶节点
		stack = stack[:len(stack)-1] // 栈顶出栈
		res = append(res, node.Data) // 将节点值加入结果数组
		node = node.Right            // 遍历右子树
	}
	return res
}

// 后序遍历（非递归）
func postorderTraversal(root *Node) []int {
	if root == nil {
		return []int{}
	}
	stack, res := []*Node{}, []int{}
	p, h := root, new(Node)
	for p != nil || len(stack) > 0 { // 当节点不为空或栈不为空时循环
		if p != nil { // 将所有左子节点入栈
			stack = append(stack, p)
			p = p.Left
		} else { // 到达叶子节点时开始出栈
			node := stack[len(stack)-1]               // 取出栈顶节点
			if node.Right != nil && node.Right != h { // 右子节点不为空且未被访问时，遍历右子树
				p = node.Right
			} else { // 右子节点为空或已被访问时，将节点值加入结果数组，标记为已访问，并出栈
				res = append(res, node.Data)
				h = node
				stack = stack[:len(stack)-1]
			}
		}
	}
	return res
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

	fmt.Println("前序遍历（非递归）:")
	fmt.Println(preorderTraversal(root))
	fmt.Println()

	fmt.Println("中序遍历（非递归）:")
	fmt.Println(inorderTraversal(root))
	fmt.Println()

	fmt.Println("后序遍历（非递归）:")
	fmt.Println(postorderTraversal(root))
	fmt.Println()

}
