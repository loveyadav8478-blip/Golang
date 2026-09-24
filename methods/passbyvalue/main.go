package main

import "fmt"

func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
	return getBillForMonth(costPerSend, numLastMonth)-getBillForMonth(costPerSend, numThisMonth)
}

func getBillForMonth(costPerSend, messagesSent int) int{
	return costPerSend * messagesSent
}

func main(){
	costPerSend := 100
	numLastMonth := 1000
	numThisMonth := 100
	fmt.Println(monthlyBillIncrease(costPerSend,numLastMonth,numThisMonth))
}