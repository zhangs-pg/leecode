package main

import (
	"fmt"
)

func is(a byte) bool {
	if (a >= 65 && a <= 90) || (a >= 97 && a <= 122) {
		return true
	}
	return false
}

func reverseOnlyLetters(S string) string {
	b := []byte(S)
	l := len(b)
	i, j := 0, l-1
	for i < j {
		if is(b[i]) == false {
			i++
		}
		if is(b[j]) == false {
			j--
		}

		if is(b[i]) == true && is(b[j]) == true {
			b[i], b[j] = b[j], b[i]
			i++
			j--
		}

	}
	return string(b)
}

func main() {
	fmt.Println(reverseOnlyLetters("-S2,_"))
}
