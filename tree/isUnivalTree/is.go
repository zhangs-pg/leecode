package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func unival(val int, node *TreeNode) bool {
	if node.Left != nil {
		if node.Left.Val == val {
			if unival(val, node.Left) == false {
				return false
			}
		} else {
			return false
		}
	}
	if node.Right != nil {
		if node.Right.Val == val {
			if unival(val, node.Right) == false {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func isUnivalTree(root *TreeNode) bool {
	val := root.Val
	return unival(val, root)
}

func New(index int, data []int) *TreeNode {
	node := &TreeNode{data[index], nil, nil}
	// 设置完全二叉树左节点  其特征是深度 *2+1为左节点  +2为右节点
	if index < len(data) && 2*index+1 < len(data) {
		node.Left = New(index*2+1, data)
	}
	if index < len(data) && 2*index+2 < len(data) {
		node.Right = New(index*2+2, data)
	}
	return node
}

func main() {
	node := New(0, []int{2, 2, 2, 5, 2})
	fmt.Println(isUnivalTree(node))
}
