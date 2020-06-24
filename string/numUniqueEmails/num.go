package main

import (
	"fmt"
	"strings"
)

func getEmail(s string) string {
	var ans string
	kv := strings.Split(s, "@")
	l := len(kv[0])
	for i := 0; i < l; i++ {

		fmt.Println(string(kv[0][i]))

		if string(kv[0][i]) == "." {
			continue
		}

		if string(kv[0][i]) == "+" {
			break
		}

		ans += string(kv[0][i])
	}
	return ans + "@" + kv[1]
}

func numUniqueEmails(emails []string) int {
	m := make(map[string]int, len(emails))
	for _, e := range emails {
		m[getEmail(e)] = 1
	}
	fmt.Println(m)
	return len(m)
}

func main() {
	a := []string{"test.email+alex@leetcode.com", "test.email.leet+alex@code.com"}
	fmt.Println(numUniqueEmails(a))
}
