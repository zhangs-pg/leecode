package main

import (
	"fmt"
	"sort"
)

type KthLargest struct {
	k    int
	nums []int
}

func Constructor(k int, nums []int) KthLargest {
	if k < len(nums) {
		sort.Ints(nums)
		nums = nums[len(nums)-k:]
	}

	buildMinHeap(nums, len(nums))
	return KthLargest{
		k:    k,
		nums: nums,
	}
}

func (this *KthLargest) Add(val int) int {
	if this.k > len(this.nums) {
		this.nums = append(this.nums, val)
		buildMinHeap(this.nums, len(this.nums))
	} else {
		if val > this.nums[0] {
			this.nums[0] = val
			heapify(this.nums, 0, len(this.nums))
		}
	}
	return this.nums[0]
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

func main() {
	arr := []int{4, 5, 8, 2}

	k := Constructor(3, arr)
	fmt.Println(k)
	for _, n := range []int{3, 5, 10, 9, 4} {
		v := k.Add(n)
		fmt.Println("result", v)
	}
}
