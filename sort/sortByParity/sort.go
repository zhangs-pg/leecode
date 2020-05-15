package main

import "fmt"

func sortArrayByParityII(A []int) []int {
	i := 0
	j := 1
	l := len(A)
	for i < l && j < l {
		for i < l && A[i]%2 == 0 {
			i += 2
		}

		for j < l && A[j]%2 != 0 {
			j += 2
		}

		if j < l && i < l {
			A[j], A[i] = A[i], A[j]
		}
	}
	return A
}

func main() {
	a := []int{4, 2, 5, 7}
	fmt.Println(sortArrayByParityII(a))
}
