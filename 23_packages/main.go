package main

import (
	// "fmt"

	"github.com/fatih/color"
	"github.com/jk872001/authApp/auth"
	"github.com/jk872001/authApp/user"
)

func main(){
	auth.LoginWithCredentials("jk","231")

	user := user.Users{
		Email : "jk@gmail.com",
		Name : "Jitesh",
	}

	// fmt.Println(user)
	color.Red(user.Name)
}
