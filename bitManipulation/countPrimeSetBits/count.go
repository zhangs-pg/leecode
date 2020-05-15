package main

import (
	"fmt"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func countPrimeSetBits(L int, R int) int {
	count := 0
	tmp := 0
	ans := 0
	for i := L; i <= R; i++ {
		tmp = i
		count = 0
		for tmp > 0 {
			if tmp&1 == 1 {
				count++
			}

			tmp >>= 1
		}
		if isPrime(count) {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(countPrimeSetBits(10, 15))
}
