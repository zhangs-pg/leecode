package main

import (
	"fmt"
	"strings"
	"unicode"
)

func letterCasePermutation(S string) []string {
	ans := [][]string{[]string{}}
	for _, s := range S {
		len := len(ans)
		if unicode.IsLetter(s) {
			fmt.Println(ans, strings.ToLower(string(s)), strings.ToUpper(string(s)))
			for j := 0; j < len; j++ {
				ans = append(ans, ans[j][:])
				ans[j] = append(ans[j], strings.ToLower(string(s)))
				ans[len+j] = append(ans[len+j], strings.ToUpper(string(s)))
			}
		} else {
			for a := 0; a < len; a++ {
				ans[a] = append(ans[a], string(s))
			}
		}
	}

	re := []string{}
	fmt.Println(ans)

	for _, x := range ans {
		re = append(re, strings.Join(x, ""))
	}
	return re
}

func main() {
	fmt.Println(letterCasePermutation("8gdG"))
}
