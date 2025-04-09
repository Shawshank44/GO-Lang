package main

import (
	"fmt"
	"packer/auth"
	"packer/user"
)

func main() {

	auth.LoginWithCredentials("Shashank", "12345")
	fmt.Println(auth.Getsession())

	user := user.User{
		Email:    "xyz@email.com",
		Username: "shashank",
	}

	fmt.Println(user)

}
