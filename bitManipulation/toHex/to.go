package main

import "fmt"

func toHex(num int64) string {
	if num < 0 {
		num += 4294967296
	}
	if num == 0 {
		return "0"
	}
	res := ""
	hex := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	for num != 0 {
		res += hex[num%16]
		num /= 16
	}
	return res
}

func main() {
	fmt.Println(toHex(26))
	fmt.Println(toHex(1))
	fmt.Println(toHex(-1))
}
