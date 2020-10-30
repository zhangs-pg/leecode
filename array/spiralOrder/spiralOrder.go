package main

import "fmt"

func getA(k int) int {
	return k - 1
}

func getB(n, k int) int {
	return n - k
}

func spiralOrder(matrix [][]int) []int {
	rst := []int{}
	m := len(matrix)
	if m == 0 {
		return rst
	}
	n := len(matrix[0])
	if m == 1 {
		return matrix[0]
	} else if n == 1 {
		for _, x := range matrix {
			rst = append(rst, x[0])
		}
	}
	var c int
	if m > n {
		c = n
	} else {
		c = m
	}
	for k := 1; k < c; k++ {
		a := getA(k)

		if len(rst) == n*m {
			return rst
		}
		for y := a; y <= getB(n, k); y++ {
			rst = append(rst, matrix[a][y])
		}

		if len(rst) == n*m {
			return rst
		}
		for y := a + 1; y <= getB(m, k); y++ {
			rst = append(rst, matrix[y][n-k])
		}

		if len(rst) == n*m {
			return rst
		}
		for y := getB(n, k) - 1; y >= a; y-- {
			rst = append(rst, matrix[m-k][y])
		}

		if len(rst) == n*m {
			return rst
		}
		for y := getB(m, k) - 1; y > a; y-- {
			rst = append(rst, matrix[y][a])
		}

	}
	return rst
}

func main() {
	a := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	fmt.Println(spiralOrder(a))
}
