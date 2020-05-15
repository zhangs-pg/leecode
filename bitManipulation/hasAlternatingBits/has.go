package main

import (
	"fmt"
)

func hasAlternatingBits(n int) bool {
	if n < 5 {
		return false
	}
	x := n & 1
	for n != 0 {
		n >>= 1
		if x == 1 {
			x = 0
		} else if x == 0 {
			x = 1
		}
		if n&1 != x {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(hasAlternatingBits(7))
}
