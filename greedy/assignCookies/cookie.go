package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	fmt.Println(g)
	glen := len(g)
	slen := len(s)
	i := 0
	j := 0
	ans := 0
	for i < glen && j < slen {
		if g[i] <= s[j] {
			i++
			j++
			ans++
		} else {
			j++
		}
	}
	return ans
}

func main() {
	findContentChildren([]int{10, 9, 8, 7}, []int{5, 6, 7, 8})
}
