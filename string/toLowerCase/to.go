package main

import (
	"fmt"
)

func toLowerCase(str string) string {
	var ans string
	for _, i := range str {

		if i >= 65 && i <= 90 {
			ans += string((i + 32))

		} else {
			ans += string(i)
		}

	}
	return ans
}

func main() {
	fmt.Println(toLowerCase("al&phaBET"))
}
