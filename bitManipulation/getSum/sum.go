package main

import "fmt"

func getSum(a int, b int) int {
	low := 0
	ca := 0
	for {
		low = a ^ b
		ca = a & b
		if ca == 0 {
			return low
		}
		a = low
		b = ca << 1
	}
	return 0
}
func main() {
	fmt.Println(getSum(2, 3))
}
