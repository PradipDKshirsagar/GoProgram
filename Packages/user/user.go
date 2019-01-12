package user

import ("fmt"
		"GoProgram/Packages/company"
		)
func UserInfo(){
	fmt.Println("I am in user packager UserInfo function")
	company.CompanyInfo();
}
