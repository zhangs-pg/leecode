package main

import "fmt"

func reverseBits(num uint32) uint32 {
	var ans uint32
	ans = 0
	for i := 31; i >= 0; i-- {
		ans = ans | ((num >> (31 - i) & 1) << i)
	}
	return ans
}

func main() {
	fmt.Println(reverseBits(43261596))
}
