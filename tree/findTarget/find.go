package main

import "fmt"

var m = make(map[int]int)

func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	_, ok := m[root.Val]
	if ok {
		return true
	} else {
		m[k-root.Val] = root.Val
	}
	return findTarget(root.Left, k) || findTarget(root.Right, k)
}

func main() {
	fmt.Println(1)
}
