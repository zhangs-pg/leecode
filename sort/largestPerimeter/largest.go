package main

import (
	"fmt"
	"sort"
)

func largestPerimeter(A []int) int {
	sort.Ints(A)
	l := len(A)
	for i := l - 3; i >= 0; i-- {
		if A[i]+A[i+1] > A[i+2] {
			return A[i] + A[i+1] + A[i+2]
		}
	}
	return 0
}

func main() {
	fmt.Println("vim-go")
}
