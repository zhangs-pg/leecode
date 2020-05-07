package main

import (
	"fmt"
)

func twoCitySchedCost(costs [][]int) int {
	for i := 0; i < len(costs); i++ {
		for j := i + 1; j < len(costs); j++ {
			if (costs[i][0] - costs[i][1]) > (costs[j][0] - costs[j][1]) {
				costs[i], costs[j] = costs[j], costs[i]
			}
		}
	}

	sum := 0
	for k := 0; k < len(costs)/2; k++ {
		sum += costs[k][0] + costs[len(costs)/2+k][1]
	}
	return sum
}

func main() {
	a := [][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}}

	fmt.Println(twoCitySchedCost(a))
}
