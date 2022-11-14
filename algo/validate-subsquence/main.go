package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	seq := []int{4, 5, 6}
	fmt.Println("[1,2,3,4,5,6,7], [4,5,6]", isSubsequence(arr, seq))
}

func isSubsequence(array []int, sequence []int) bool {
	seqID := 0
	for _, val := range array {
		if seqID == len(sequence) {
			break
		}
		if val == sequence[seqID] {
			seqID++
		}
	}
	return seqID == len(sequence)
}
