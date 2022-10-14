package main
import "fmt"

func changeVal(myptr *int) {
	*myptr = 12
}

func main() {
	val := 10
	fmt.Println("val is:", val)
	changeVal(&val)
	fmt.Println("val now is:", val)
}
