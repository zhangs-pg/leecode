package main

import "fmt"

func lastStoneWeight(stones []int) int {
	for len(stones) < 3 {
		stones = append(stones, 0)
	}
	buildMaxHeap(stones, len(stones))
	for stones[2] != 0 || stones[1] != 0 {
		if stones[1] > stones[2] {
			stones[0] = stones[0] - stones[1]
			stones[1] = 0
			heapify(stones, 1, len(stones))
		} else {
			stones[0] = stones[0] - stones[2]
			stones[2] = 0
			heapify(stones, 2, len(stones))
		}
		heapify(stones, 0, len(stones))
	}
	return stones[0]
}

func buildMaxHeap(arr []int, arrLen int) {
	for i := arrLen / 2; i >= 0; i-- {
		heapify(arr, i, arrLen)
	}
}

func heapify(arr []int, i, arrLen int) {
	left := i*2 + 1
	right := i*2 + 2
	largest := i
	if left < arrLen && arr[left] > arr[largest] {
		largest = left
	}
	if right < arrLen && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		swap(arr, i, largest)
		heapify(arr, largest, arrLen)
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	arr := []int{2, 7, 4, 1, 8, 1}
	ans := lastStoneWeight(arr)
	fmt.Println(ans)
}
