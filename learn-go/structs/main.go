package main

import (
	"fmt"
)

type customer struct {
	name  string
	money float64
}
type rectangle struct {
	length, height float64
}

func getCustomer(c customer) {
	fmt.Println(c)
}

func (r rectangle) Area() float64 {
	return r.length * r.height
}

func main() {
	var customer customer
	customer.name = "yo"
	customer.money = 5.000
	getCustomer(customer)
	rectangle := rectangle{10.5, 5.5}
	fmt.Println("rectangle area:",rectangle.Area())
}
