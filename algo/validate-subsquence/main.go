package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 3, 4, 5, 1}
	seq := []int{2, 3, 4}
	fmt.Println("[2,3,4,5,1], [2,3,4]", isSubsequence(arr, seq))
}

func isSubsequence(array []int, sequence []int) bool {
	arrID := 0
	seqID := 0
	for arrID < len(array) && seqID < len(sequence) {
		if array[arrID] == sequence[seqID] {
			seqID += 1
		}
		arrID += 1
	}
	return arrID == seqID
}
