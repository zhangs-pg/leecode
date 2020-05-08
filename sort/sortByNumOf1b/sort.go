package main

import (
	"fmt"
	"sort"
)

func getCount(i int) int {
	count := 0
	for i > 0 {
		if i&1 == 1 {
			count++
		}
		i >>= 1
	}
	return count
}

func sortByBits(arr []int) []int {
	m := make(map[int]int, len(arr))

	for _, i := range arr {
		m[i] = getCount(i)
	}

	sort.Slice(arr, func(i, j int) bool {
		if m[arr[i]] == m[arr[j]] {
			return arr[i] <= arr[j]
		}
		return m[arr[i]] <= m[arr[j]]
	})
	return arr
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(sortByBits(a))
}
