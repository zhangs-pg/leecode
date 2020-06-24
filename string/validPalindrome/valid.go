package main

import "fmt"

func isPalindrome(s string) bool {
	l := len(s)
	i := 0
	j := l - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func validPalindrome(s string) bool {
	l := len(s)
	i := 0
	j := l - 1
	for i < j && j >= 0 && i < l-1 {
		if s[i] != s[j] {
			return isPalindrome(s[i+1:j+1]) || isPalindrome(s[i:j])
		}
		i += 1
		j -= 1
	}
	return true
}

func main() {
	fmt.Println(validPalindrome("lcuaucul"))
	s := "123"
	for _, x := range s {
		fmt.Println(x)
	}
}
