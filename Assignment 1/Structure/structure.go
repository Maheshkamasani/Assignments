package main

import (  
    "fmt"
)

type User struct {  
    firstName string
    lastName  string
    age       int
    salary    int
}

func main() {

    //creating struct specifying field names
    user1 := User{
        firstName: "Maheswar Reddy",
        age:       22,
        salary:    10000,
        lastName:  "Kamasani",
    }

    //creating struct without specifying field names
    user2 := User{"Chinnapa Reddy", "Kanakanti", 22, 25000}

    fmt.Println("User 1", user1)
    fmt.Println("User 2", user2)
}