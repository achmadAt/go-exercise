package main

import (
	"fmt"
)

var print = fmt.Println

type contact struct {
	name   string
	number int
}
type business struct {
	company string
	address string
	contact contact
}

func main() {
	contact := contact{
		"human",
		123,
	}
	business := business{
		"earth",
		"galaxy",
		contact,
	}
	print(business)
}
