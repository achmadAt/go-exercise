package main
import (
	"fmt"
)

var print = fmt.Println
func nums(channel chan int){
	channel <- 1
	channel <- 2
	channel <- 3
}

func main() {
	number := make(chan int)
	go nums(number)
	for i:= 0; i< 3; i++ {
		print(<-number)
	}
}
