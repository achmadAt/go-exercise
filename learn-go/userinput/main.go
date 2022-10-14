package main
import (
	"fmt"
	"bufio"
	"os"
	"math/rand"
)

var print = fmt.Println
func randomId(length int) int {
	return rand.Intn(length-1) + 1
} 
func main() {
	print("insert your name")
	greets := []string{"hy","hello","yo"}
	greet := greets[randomId(len(greets))]
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		print(err)
	}
	print(greet + "! " + name)
}
