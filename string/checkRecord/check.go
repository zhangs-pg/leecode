package main

import (
	"fmt"
	"strings"
)

func checkRecord(s string) bool {
	acount := strings.Count(s, "A")
	if acount > 1 {
		return false
	}
	l := strings.Split(s, "A")
	for _, x := range l {
		x2 := strings.Split(x, "P")
		for _, i := range x2 {
			if len(i) > 2 {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println(checkRecord("PPALLP"))
}
