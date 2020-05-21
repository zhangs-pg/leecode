package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

func check(n1 *TreeNode, n2 *TreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	}
	if n1 == nil || n2 == nil {
		return false
	}
	if n1.Val == n2.Val {
		return check(n1.Left, n2.Left) && check(n1.Right, n2.Right)
	}
	return false
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	if s.Val == t.Val {
		if check(s, t) == true {
			return true
		}
	}
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func main() {
	node := New(0, []int{3, 4, 5, 1, 2})
	node2 := New(0, []int{4, 1, 0, 0, 2})

	fmt.Println(isSubtree(node, node2))
}
