package main

import "fmt"

func main(){
	var smsSendingLimits int
	var costPerSms float64
	var hasPermission bool
	var username string

	fmt.Printf(
		"%v %f %v %q\n",
		smsSendingLimits,
		costPerSms,
		hasPermission,
		username,
	)
}