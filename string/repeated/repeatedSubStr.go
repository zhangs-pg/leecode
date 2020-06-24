package main

import (
	"fmt"
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	tmp := ""
	for i, _ := range s {
		tmp += string(s[i])
		if strings.Count(s, tmp)*len(tmp) == len(s) && i != len(s)-1 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(repeatedSubstringPattern("ababc"))
}
