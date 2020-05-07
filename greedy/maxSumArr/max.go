package main

import (
	"fmt"
)

func sum(arr []int) int {
	sum := 0
	for _, i := range arr {
		sum += i
	}
	return sum
}
func buildMinHeap(arr []int, arrLen int) {
	for i := arrLen / 2; i >= 0; i-- {
		heapify(arr, i, arrLen)
	}
}

func heapify(arr []int, i, arrLen int) {
	left := i*2 + 1
	right := i*2 + 2
	largest := i
	if left < arrLen && arr[left] < arr[largest] {
		largest = left
	}
	if right < arrLen && arr[right] < arr[largest] {
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

func largestSumAfterKNegations(A []int, K int) int {
	buildMinHeap(A, len(A))
	if K > 0 {
		for K > 0 {
			A[0] = -A[0]
			heapify(A, 0, len(A))
			K--
		}
	}
	return sum(A)
}

func main() {
	fmt.Println(largestSumAfterKNegations([]int{-2, 9, 9, 8, 4}, 5))
}
