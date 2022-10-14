package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)
	randNum := rand.Intn(5) + 1
	for true {
		fmt.Print("guees num between 0 to 5:")
		fmt.Println("random num is:", randNum)
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		guess = strings.TrimSpace(guess)
		iguess, err := strconv.Atoi(guess)
		if err != nil {
			panic(err)
		}
		if iguess > randNum {
			fmt.Println("too big")
		}
		if iguess < randNum {
			fmt.Println("too small")
		}
		if iguess == randNum {
			fmt.Println("pass")
			break
		}
	}
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	i := 0
	for i < len(nums) {
		fmt.Println(i)
		i++
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
}
