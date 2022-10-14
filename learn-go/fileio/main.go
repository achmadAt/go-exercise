package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Create("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	iprime := []int{2, 3, 5, 7, 11}
	var sprime []string
	for _, i := range iprime {
		sprime = append(sprime, strconv.Itoa(i))
	}
	for _, num := range sprime {
		_, err := file.WriteString(num + "\n")
		if err != nil {
			fmt.Println(err)
		}
	}
	file, err = os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		fmt.Println("prime :", scan.Text())
	}
	if err := scan.Err(); err != nil {
		fmt.Println(err)
	}

}
