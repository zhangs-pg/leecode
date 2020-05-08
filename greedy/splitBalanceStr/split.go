package main

import (
	"fmt"
)

//stack
//func balancedStringSplit(s string) int {
//	stack := ""
//	res := 0
//	for _, i := range []byte(s) {
//		stack += string(i)
//		c_1 := strings.Count(stack, "L")
//		c_2 := strings.Count(stack, "R")
//		if c_1 == c_2 && c_1 > 0 {
//			stack = ""
//			res++
//		}
//	}
//	return res
//}

func balancedStringSplit(s string) int {
	count := 0
	balance := 0
	for _, i := range s {
		if string(i) == "L" {
			balance--
		} else if string(i) == "R" {
			balance++
		}
		if balance == 0 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(balancedStringSplit("RLRRLLRLRL"))
}
