package main

import (
	"fmt"
	"unicode"
)

func isDig(s string) bool {
	return unicode.IsDigit(s)
}

func isLet(s string) bool {
	return unicode.IsLetter(s)
}

func reformat(s string) string {
	l := len(s)
	i := 0
	j := 1
	ans := ""
	for {
		if isDig(string(s[i])) == true {
			for isLet(string(s[j])) == false {
				j++
			}
			ans += string(s[i])
			ans += string(s[j])
		} else if isLet(string(s[i])) == true {
			for isDig(string(s[j])) == false {
				j++
			}
			ans += string(s[i])
			ans += string(s[j])
		}
	}
}

func main() {
	fmt.Println("vim-go")
}
