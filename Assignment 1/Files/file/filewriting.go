package main

import ("fmt"
		"os"
		"bufio"
		"io/ioutil"
)



func main ()  {
	message := []byte("Hello Friends")
	err := ioutil.WriteFile("new.txt", message,0644)
	if err != nil{
		fmt.Println(err)
	}

  // creating a file

  f,err := os.Create("mahesh.txt")
  defer f.Close()
  if err != nil{
  	fmt.Println(err)
  }
  f.WriteString("Hello friends !!!")
  f.Sync()

  w := bufio.NewWriter(f)

  w.WriteString("Welcome to Kloudone.")

  w.Flush()
}