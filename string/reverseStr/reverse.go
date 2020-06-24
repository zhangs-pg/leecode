package main

import (
	"fmt"
	"math"
)

func reverseStr(s string, k int) string {
	str := []byte(s)
	ln := len(s)
	for st := 0; st < ln; st += 2 * k {
		i := st
		j := int(math.Min(float64(st+k-1), float64(ln-1)))
		for i < j {
			str[i], str[j] = str[j], str[i]
			i++
			j--
		}
	}
	return string(str)
}

func main() {
	fmt.Println(reverseStr("abcde", 2))
}
