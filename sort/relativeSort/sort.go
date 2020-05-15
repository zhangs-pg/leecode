package main

import "fmt"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	a := make([]int, 1001)
	ans := []int{}
	for _, i := range arr1 {
		a[i] += 1
	}
	for _, j := range arr2 {
		for a[j] > 0 {
			ans = append(ans, j)
			a[j] -= 1
		}
	}
	for k := 0; k < 1001; k++ {
		for a[k] > 0 {
			ans = append(ans, k)
			a[k]--
		}
	}
	return ans
}

func main() {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	fmt.Println(relativeSortArray(arr1, arr2))
}
