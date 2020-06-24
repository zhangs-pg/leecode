package main

import (
	"fmt"
)

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	} else {
		l := 0
		for _, c := range root.Children {
			t := maxDepth(c)
			if t > l {
				l = t
			}
		}
		return l + 1
	}
}

func main() {
	fmt.Println("vim-go")
}
