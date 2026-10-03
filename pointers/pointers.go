package main

import "fmt"

func main() {
	num := 10
	fmt.Println("Before change num ", num)
	changeNum(&num)
	fmt.Println("After change num ", num)
}

func changeNum(num *int) {
	*num = 20
}
