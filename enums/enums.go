package main

import "fmt"

type OrderStatus int

const(
	Recieved OrderStatus = iota
	Pending
	Confirmed
	Shipped
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to ",status)
}

func main() {
	changeOrderStatus(Recieved)
}