package main

import "fmt"

type Address struct {
	User
	City string
	State string
	Email string
}

type User struct{
	Name string
	Phone int
	Age int
	City string
	State string
	Email string

}


func main(){
	var details User

	details.Name  = "Maheswar Reddy Kamasani"
	details.Phone = 6382839908
	details.Age   = 22
	details.City  = "Kadapa" 
	details.State = "Andhra Pradesh"
	details.Email = "maheshkamasani1729@gmail.com"

	printUser(details)
}

func printUser(userDetails User){

		fmt.Println(userDetails.Name)
		fmt.Println(userDetails.Phone)
		fmt.Println(userDetails.Age)
		fmt.Println(userDetails.City)
		fmt.Println(userDetails.State)
		fmt.Println(userDetails.Email)
}

