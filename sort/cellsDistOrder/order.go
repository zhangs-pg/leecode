package main

import (
	"fmt"
	"math"
	"sort"
)

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	m := make(map[int][][]int, R*C)
	keys := []int{}
	ans := [][]int{}
	for x := 0; x < R; x++ {
		for y := 0; y < C; y++ {
			tmp := int(math.Sqrt(float64((x-r0)*(x-r0))) + math.Sqrt(float64((y-c0)*(y-c0))))
			m[tmp] = append(m[tmp], []int{x, y})
			keys = append(keys, tmp)
		}
	}

	sort.Ints(keys)

	fmt.Println(keys, m)
	x := -1
	for _, i := range keys {
		if x == i {
			continue
		}
		for _, v := range m[i] {
			ans = append(ans, v)
		}
		x = i
	}
	return ans
}

func main() {
	fmt.Println(allCellsDistOrder(3, 5, 2, 3))
}
