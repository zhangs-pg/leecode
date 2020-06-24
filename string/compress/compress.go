package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	l := len(chars)
	i := 0
	j := 1
	if l == 1 {
		return 1
	}

	ans := chars[0:0]

	for i < l {
		if chars[j] == chars[i] {
			j++
			if j >= l {
				ans = append(ans, chars[i])
				if j-i > 1 {
					for _, n := range []byte(strconv.Itoa(j - i)) {
						ans = append(ans, n)
					}
				}
				break
			}
		} else {
			ans = append(ans, chars[i])
			if j-i > 1 {
				for _, n := range []byte(strconv.Itoa(j - i)) {
					ans = append(ans, n)
				}
			}

			i = j
		}
	}

	return len(ans)
}

func main() {
	x := "aabbbcccc"
	a := []byte(x)
	fmt.Println(compress(a))
}
