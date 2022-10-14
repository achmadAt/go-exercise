package main

import (
	"fmt"
)

var print = fmt.Println

func useFunc(f func(int, int) int, x, y int) {
	print("answer :", (f(x, y)))
}
func sumVal(x, y int) int {
	return x + y
}
func main() {
	intSum := func(x, y int) int { return x + y }
	print("5+4=", intSum(5, 4))
	useFunc(sumVal, 5, 8)
}
