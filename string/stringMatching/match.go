package main

import (
	"fmt"
	"strings"
)

func stringMatching(words []string) []string {
	l := len(words)
	s := ""
	ans := []string{}
	for i := 0; i < l; i++ {
		s += words[i] + ","
	}

	for i := 0; i < l; i++ {
		if strings.Count(s, words[i]) > 1 {
			ans = append(ans, words[i])
		}
	}
	return ans
}

func main() {
	fmt.Println(stringMatching([]string{"mass", "as", "hero", "superhero"}))
}
