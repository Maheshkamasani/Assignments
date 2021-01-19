package main

import "fmt"

type Address struct {
    city    string
    state string
    country string
}

type User struct {
    name    string
    age     int
    address Address
}

func main() {

    p := User{
        name: "Maheswar Reddy Kamasani",
        age:  22,
        address: Address{
            city:    "Kadapa",
	    state: "Andhra Pradesh",
            country: "India",
        },
    }

    fmt.Println("Name:", p.name)
    fmt.Println("Age:", p.age)
    fmt.Println("City:", p.address.city)
    fmt.Println("State:", p.address.state)
    fmt.Println("Country:", p.address.country)
}