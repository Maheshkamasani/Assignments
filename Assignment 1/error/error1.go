package main

import "fmt"

func main(){
	defer catchIssues("divide")  
	divide()
	fmt.Println("End of the Main")
}

func catchIssues(functionName string){
	rec := recover()
	if rec != nil{
		fmt.Println("somthing has gone  worng",functionName)
		fmt.Println("Reson:",rec)
	}
}

func divide(){

	num := 20
	den := 0
	fmt.Println("Enter a number")
	fmt.Scanf("%d",&den)

	if den == 0 {
		panic("Oops !!! Cannot Enter 0")
	}else{
		div := num/den
		fmt.Println(div)
	}
}




