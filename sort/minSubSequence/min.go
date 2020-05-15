package main

import (
	"fmt"
	"sort"
)

func minSubsequence(nums []int) []int {
	l := len(nums)
	if l == 1 {
		return nums
	}
	sort.Ints(nums)

	sum := 0
	tmp := 0

	ans := []int{}

	for _, v := range nums {
		sum += v
	}
	for i := l - 1; i >= 0; i-- {
		tmp += nums[i]
		ans = append(ans, nums[i])
		if tmp > sum-tmp {
			return ans
		}
	}
	return []int{}
}

func main() {
	fmt.Println("vim-go")
}
