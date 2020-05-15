package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	p := head
	l := []int{}
	for p.Next != nil {
		l = append(l, p.Val)
		p = p.Next
	}
	n := len(l)
	ans := 0
	for k, v := range l {
		ans += int(math.Pow(2, float64(n-k-1))) * v
	}
	return ans
}

func main() {
	fmt.Println()
}
