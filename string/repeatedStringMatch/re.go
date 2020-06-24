package main

import (
	"fmt"
	"strings"
)

func repeatedStringMatch(A string, B string) int {
	tmp := A
	i := 0
	for {
		x := strings.Count(tmp, B)
		if x > 0 {
			return i + 1
		} else if len(tmp) > len(A)+len(B) {
			return 0
		} else {
			tmp += A
			i++
		}
	}
	return 0
}

func main() {
	fmt.Println(repeatedStringMatch("aaaaaaaaaaaaaaaaaaaaaab", "ba"))
}
