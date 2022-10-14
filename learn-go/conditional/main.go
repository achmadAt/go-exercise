package main
import (
	"fmt"
	"math/rand"
)
func randomNum() int {
	return rand.Intn(10-1)+1
}
func main() {
	if randomNum() / 2 == 0 {
		fmt.Println("hello")
	} else {
		fmt.Println("hy")
	}
}
