package main

import "fmt"

func lemonadeChange(bills []int) bool {
	c5 := 0
	c10 := 0
	c20 := 0
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			c5++
		} else if bills[i] == 10 {
			c5--
			c10++
			if c5 < 0 || c10 < 0 || c20 < 0 {
				fmt.Println(1, i, c5, c10, c20)

				return false
			}
		} else if bills[i] == 20 {
			c5--
			c10--
			c20++

			if c10 < 0 {
				c10++
				c5--
				c5--
			}
			if c5 < 0 || c10 < 0 {
				fmt.Println(2, i, c5, c10, c20)
				return false
			}

		}
	}

	return true
}

func main() {
	a := []int{5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5}
	fmt.Println(lemonadeChange(a))
}
