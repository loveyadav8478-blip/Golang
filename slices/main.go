package main

import (
	"fmt"
	"slices"
)

func main() {
	var num1 = []int{1,2,3}
	var num2 = []int{1,22,3}
	num1[0] = 1
	num2[0] = 1

	fmt.Println(slices.Equal(num1,num2))
}