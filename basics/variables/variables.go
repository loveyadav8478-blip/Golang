package main

import "fmt"

func Add(a int, b int) int{
	return a+b
}

func main(){
	var smsSendingLimits int
	var costPerSms float64
	var hasPermission bool
	var username string

	var nums [6]int
	var dNums [] int
	for i:= 0; i<=100; i++{
		dNums = append(dNums, i);
	}
	fmt.Print(dNums)
	println()
	for i := 0; i<len(nums); i++{
		nums[i] = i
	}
	for i := len(nums)-1; i>=0; i--{
		fmt.Print(nums[i]," ")
	}
	println()
	

	fmt.Printf(
		"%v %f %v %q\n",
		smsSendingLimits,
		costPerSms,
		hasPermission,
		username,
	)
	fmt.Print(Add(3,4))
}