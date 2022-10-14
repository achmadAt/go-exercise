package main
  
import "fmt"
  
// The addition function returns
// the sum of the elements
// Unexported function
func addition(val ...int) int {
    s := 0
  
    for x := range val {
        s += val[x]
    }
  
    fmt.Println("Total Sum: ", s)
    return s
}
  
// Main function
func main() {
  
    addition(23, 546, 65, 42, 21, 24, 67)
}
