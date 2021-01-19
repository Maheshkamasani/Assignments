package main

import(
	"fmt"
	"time"
)


func sayHello(s string){
	for i :=0; i<3; i++{
		fmt.Println(s)
		time.Sleep(100 * time.Millisecond)
	}
}

func main(){
	go sayHello("Hi")
	go sayHello("Mahesh")
	go sayHello("Bye")
	time.Sleep(time.Second)
fmt.Println("done")
	fmt.Println("Success")
}


