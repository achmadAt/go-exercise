package main

import "fmt"

func main() {
	a := [4]int{1, 2, 3, 4}
	fmt.Println(len(a))
	for i:=0;i<len(a);i++{
		fmt.Println(a[i])
	}
}
