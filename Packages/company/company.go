package company

import "fmt"

func CompanyInfo(){
	fmt.Println("I am in company packager CompanyInfo function")
}

func GetName(){
	fmt.Println("In GetName of company")
	//user.UserInfo()   //cycle is created so avoid that condition by design good use-cases
}