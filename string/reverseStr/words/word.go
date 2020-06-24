package main

import (
	"fmt"
)

func reverseWords(s string) string {
	l := []byte(s)
	left := 0
	for i, v := range s {
		if v == " " || i == len(s)-1 {
			r := i - 1
			if i == len(s)-1 {
				r = i
			}
			for left < r {
				l[left], l[r] = l[r], l[left]
				left++
				r--
			}
			left = i + 1
		}
	}
	return string(l)
}

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}
