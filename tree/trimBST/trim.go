package main

import "fmt"

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > R {
		return trimBST(root.Left)
	} else if root.Val < L {
		return trimBST(root.Right)
	} else {
		root.Left = trimBST(root.Left, L, R)
		root.Right = trimBST(root.Right, L, R)
		return root
	}
}

func main() {
	a := []int{1, 2, 3}
	a = append(a[0 : len(a)-1])
	fmt.Println(a)
}
