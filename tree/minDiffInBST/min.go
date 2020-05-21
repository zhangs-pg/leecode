package main

import (
	"fmt"
	"math"
)

func cal(l []int, v int) int {
	var a float64
	a = 100000000.0
	for _, i := range l {
		a = math.Min(math.Abs(float64(v-i)), a)
	}
	return int(a)
}

func dfs(node *TreeNode, l *[]int, a *int) {
	if node != nil {
		*a = cal(*l, node.Val)
		*l = append(*l, node.Val)
		dfs(node.Left, l, a)
		dfs(node.Right, l, a)
	}
}

func minDiffInBST(root *TreeNode) int {
	var ans int
	ans = 0
	l := []int{}
	dfs(root, &l, &ans)
	return ans
}

func main() {
	fmt.Println("vim-go")
}
