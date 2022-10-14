package main
import (
	"fmt"
)

func sumOfNums(slices []int) int {
	sum := 0
	for _, num := range slices {
		sum += num
	}
	return sum
}

func main() {
	nums := []int{1,2,3,4,5,6,7,8,9}
	arr := [5]int{1,2,3}
	fmt.Println(sumOfNums(nums))
	fmt.Println(arr)
	fmt.Println(arr[0])
}
