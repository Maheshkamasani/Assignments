 package main

 import "fmt"
 func main(){

 	myChannel := make(chan string)

 	go func(){
 		myChannel <- "Maheswar Reddy Kamasani"
 		myChannel <- "Kloudone"
 	}()
 	text := <-myChannel

 	fmt.Println(text)
 	fmt.Println(<-myChannel)
}