package main

import (
	"fmt"

	"github.com/loveyadav8478-blip/Golang/user"
	"github.com/loveyadav8478-blip/Golang/auth"
)

func main() {

	auth.LoginWithUsernameAndPassword("LoveYadav","test123")

	fmt.Println(auth.GetSession())
	user := user.User{
		Name: "Love Yadav",
		Email: "user.12@gmail.com",
	}
	fmt.Println(user.Email, user.Name)

}