package main

import (
	"fmt"
	"strconv"
)

type watch struct {
	nums  []int
	visit []int
	ans   []string
	n     int
}

func handle_date(wa *watch) string {
	n_h := 0
	n_m := 0
	for i := 0; i < len(wa.nums); i++ {
		if wa.visit[i] == 0 {
			continue
		}
		if i < 4 {
			n_h += wa.nums[i]
		} else {
			n_m += wa.nums[i]
		}
	}
	result := "" + strconv.Itoa(n_h) + ":"
	if n_m < 10 {
		result += "0" + strconv.Itoa(n_m)
	} else {
		result += strconv.Itoa(n_m)
	}
	return result
}

func calc_sum(wa *watch) bool {
	n_h := 0
	n_m := 0
	for i := 0; i < len(wa.nums); i++ {
		if wa.visit[i] == 0 {
			continue
		}
		if i < 4 {
			n_h += wa.nums[i]
		} else {
			n_m += wa.nums[i]
		}
	}
	return n_h < 12 && n_m < 60
}

func helper(wa *watch, step int, start int) {
	if step == wa.n {
		wa.ans = append(wa.ans, handle_date(wa))
	} else {
		for i := start; i < len(wa.nums); i++ {
			wa.visit[i] = 1
			if calc_sum(wa) == false {
				wa.visit[i] = 0
				continue
			}
			helper(wa, step+1, i+1)
			wa.visit[i] = 0
		}
	}
}

func readBinaryWatch(num int) []string {
	wa := &watch{
		nums:  []int{8, 4, 2, 1, 32, 16, 8, 4, 2, 1},
		visit: make([]int, 10),
		n:     num,
	}
	helper(wa, 0, 0)
	return wa.ans
}

func main() {
	fmt.Println(readBinaryWatch(1))
}
