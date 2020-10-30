package main

import "fmt"

var m = make(map[int]int)
var c2 = 0
var c1 = 0

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	} else {
		var ans = 0
		v, ok := m[n-2]
		if ok {
			ans += v
		} else {
			c2 = climbStairs(n - 2)
			m[n-2] = c2
			ans += c2
		}

		v, ok = m[n-1]
		if ok {
			ans += v
		} else {
			c1 = climbStairs(n - 1)
			m[n-1] = c1
			ans += c1
		}
		return ans
	}
}

func main() {
	fmt.Println(climbStairs(7))
}
