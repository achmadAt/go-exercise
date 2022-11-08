package main

import (
	"fmt"
)

func main() {
	a := []int{3, 2, 2, 3}
	fmt.Println("jumlah nomor yang ada di array [3,2,2,3] tidak sama dengan 3", removeElement(a, 3))
}

func removeElement(nums []int, val int) int {
	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[idx] = nums[i]
			idx += 1
		}
	}
	return idx
}
