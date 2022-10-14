package main
import (
	"fmt"
)
func calculate(nums ...int) (string, int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return "the sum is", sum
}

func main() {
	fmt.Println(calculate(1,2,3,4,5,6,7,8,9))
}
