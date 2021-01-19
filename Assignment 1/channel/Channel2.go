package main

import "fmt"

func sentValue(details chan string){

	details <- "Maheswar Reddy Kamasani"
	details <- "Chinnapa Reddy Kanakanti"
	details <- "Babu Reddy Sirasani"
}

func main (){

	values := make(chan string)
    
	go sentValue(values)

	value :=  <- values

	defer fmt.Println(value)
		  fmt.Println(<-values)
		  fmt.Println(<-values)
}
