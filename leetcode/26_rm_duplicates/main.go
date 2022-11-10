package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 2}
	removeDuplicates(a)
}

func removeDuplicates(nums []int) int {
	length := len(nums)
	idx := 0
	if length <= 1 {
		return length
	}
	for i := 1; i < length; i++ {
		if nums[idx] != nums[i] {
			idx++
			nums[idx] = nums[i]
		}
	}

	fmt.Println(nums, "with length", length)
	return idx + 1
}
