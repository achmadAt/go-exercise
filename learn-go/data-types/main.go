package main
import (
	"fmt"
	"reflect"
)
var priln = fmt.Println

func main() {
	num := 12
	priln(reflect.TypeOf(num))
}
