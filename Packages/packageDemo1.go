package main

import ("fmt"
		"GoProgram/Packages/user"
		//"GoProgram/Packages/company"
		)


func main(){
	fmt.Println("I am in main package  function")
	user.UserInfo()

	//company.getName() // cannot refer to unexported name company.getName
}
