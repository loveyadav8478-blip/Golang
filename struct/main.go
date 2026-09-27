package main

import "fmt"


type user struct {
	name string
	number int
}

type messageToSend struct{
	msg string
	sender user
	reciever user
}

type car struct{
	make string
	height int
	capacity float32
	frontWheel wheel
	backWheel wheel
}

type wheel struct{
	radius int
	material string
}

func canSendMsg (mToSend messageToSend) bool{
	if mToSend.msg=="" || mToSend.reciever.name=="" || mToSend.reciever.number==0 || mToSend.sender.name=="" || mToSend.sender.number==0{
		return false
	}
	return true;
}

func main() {
	myCar := car{}
	fmt.Printf("Capacity of car : %f BackWheel radius : %v",myCar.capacity,myCar.backWheel.radius)
	println()
	myCar.frontWheel.radius = 5
	myCar.backWheel.radius = 5
	myCar.capacity = 100
	fmt.Printf("Capacity of car : %f BackWheel radius : %v",myCar.capacity,myCar.backWheel.radius)
	println()
	msg := messageToSend{} 
	msg.msg = "Message for user"
	msg.reciever.name = "Love"
	msg.reciever.number = 10
	msg.sender.name = "Raj"
	msg.sender.number = 20
	fmt.Println(canSendMsg(msg))
}
