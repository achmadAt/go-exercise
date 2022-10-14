package main
import "fmt"

var print = fmt.Println
func main() {
	var something map[string]int
	something = make(map[string]int)
	something["one"] = 1
	print(something)
}
