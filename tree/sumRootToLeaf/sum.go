package main

import (
	"fmt"
	"math"
)

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

func delLastNum(sum []int) []int {
	sum = sum[0 : len(sum)-1]
	return sum
}

func getAns(s []int) int {
	result := 0
	l := len(s) - 1
	for k, v := range s {
		result += int(math.Pow(2, float64(l-k))) * v
	}
	return result
}

func sumNode(node *TreeNode, s *[]int, a *int) {
	if node != nil {
		*s = append(*s, node.Val)
		if node.Left == nil && node.Right == nil {
			*a += getAns(*s)
			//fmt.Println("-", *s, reflect.TypeOf(*s), *a)
			*s = delLastNum(*s)
			//fmt.Println("+", *s, *a)
		} else {
			if node.Left != nil {
				sumNode(node.Left, s, a)
			}
			if node.Right != nil {
				sumNode(node.Right, s, a)
			}
			*s = delLastNum(*s)
		}
	}
}

func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := []int{root.Val}
	var ans int
	if root.Left != nil {
		sumNode(root.Left, &sum, &ans)
	}
	if root.Right != nil {
		sumNode(root.Right, &sum, &ans)
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	return ans
}

func main() {
	node := New(0, []int{1})
	fmt.Println(sumRootToLeaf(node))
}
