package main

import "fmt"

type paymenter interface{
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct{
	gateway paymenter
}
type razorPay struct{}
type stripe struct{}

func(p payment) makePayment(amount float32){
	p.gateway.pay(amount)
}

func(p razorPay) pay(amount float32){
	fmt.Println("Payment successful using Razorpay of ", amount)
}
func(p razorPay) refund(amount float32, account string){
	fmt.Println("Payment successful using Razorpay of account ", amount,account)
}

func(p stripe) pay(amount float32){
	fmt.Println("Payment successful using Stripe of ", amount)
}

func(p stripe) refund(amount float32, account string){
	fmt.Println("Payment successful using Razorpay of account ", amount,account)
}
func main(){

	newRazorPaymentpayment := razorPay{}
	newPayment := payment{
		gateway: newRazorPaymentpayment,
	}
	newPayment.makePayment(100)	
}