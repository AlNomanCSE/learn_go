package main

import (
	"fmt"

	"github.com/AlNomanCSE/learngo/auth"
	"github.com/AlNomanCSE/learngo/user"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWIthCredentials("noman", "n14om13@n14")
	session := auth.GeSession()
	color.Red("%s", session)
	user := user.User{
		Email: "abdullahalnomancse@gmail.com",
		Name:  "Abdullah Al Noman",
	}
	fmt.Println(user)
}
