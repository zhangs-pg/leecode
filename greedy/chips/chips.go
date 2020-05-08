package main

import "fmt"

func minCostToMoveChips(chips []int) int {
	num1 := 0
	num2 := 0
	for _, c := range chips {
		if c%2 == 1 {
			num1++
		} else {
			num2++
		}
	}
	if num1 > num2 {
		return num2
	} else {
		return num1
	}
}

func main() {
	fmt.Println("vim-go")
}
