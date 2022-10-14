package main

import "fmt"

var print = fmt.Println

type MyConstraint interface {
	int | float64
}

func getSumGen[T MyConstraint](x T, y T) T {
	return x + y
}

func main() {
	print("5 + 4 =", getSumGen(5, 4))
	print("0.5 + 0.7 =", getSumGen(0.5, 0.7))
}
