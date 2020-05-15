package main

import "fmt"

func hammingDistance(x int, y int) int {
	count := 0
	for x != 0 || y != 0 {
		if x&1 != y&1 {
			count++
		}
		x >>= 1
		y >>= 1
	}
	return count
}

func main() {
	fmt.Println(hammingDistance(1, 4))
}
