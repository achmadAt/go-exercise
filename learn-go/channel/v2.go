package main
import (
	"fmt"
	"time"
)

var print = fmt.Println

func main() {
	channel := make(chan int)
	go func(ch chan<- int, x int){
		time.Sleep(time.Second)
		ch <- x * 120
	}(channel, 9)
	done := make(chan struct{})
	go func(ch <- chan int) {
		n := <-ch
		print(n)
		time.Sleep(time.Second)
		done <- struct{}{}
	}(channel)
	<-done
	print("done")
}
