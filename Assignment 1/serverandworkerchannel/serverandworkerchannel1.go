package main

import(
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

func sayHello(s string){
	for i :=0; i<3; i++{
		fmt.Println(s)
		time.Sleep(100 * time.Millisecond)
	}
	wg.Done()
}

func main(){
	wg.Add(1)
	go sayHello("Bye")
wg.Add(1)
	go sayHello("Mahesh")
wg.Add(1)
	go sayHello("Hi")
	wg.Wait()
	fmt.Println("done")
	time.Sleep(time.Second)
	fmt.Println("Success")
}