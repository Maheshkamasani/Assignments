package main


import (
	    "fmt"
		"encoding/json"
		)	

type Preson struct{
	Name string `json:"name"`
	Phone int   `json:"phone"`
	Age int     `json:"age"`
	Gender string `json:"gender"`
	City []string `json:"city"`
}

func main () {

	user := &Preson{Name :"Maheswar Reddy Kamasani",Phone:6382839908,Age:22,Gender: "Male", City: []string{"Kadapa", "Andhra Pradesh"}}
	data, _:= json.Marshal(user)
	fmt.Println(string(data))


	 user2 := `{"Name":"Gangi Reddy Kamasani","Phone":7675903844,"Age":22,"Gender":"Male","City":["Galiveedu","Andhra Pradesh"]}`
	  data2 := &Preson{}

	  json.Unmarshal([]byte(user2),data2)
	  fmt.Println(data2)
}