package main

import (
	
	"fmt"
	"strconv"
	//"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)


	//Item is a struct that  groups  all  necessary fields  into a  single  unit
	// type Item  struct{
	// 	Data string 
	// 	OtherData int
	// }

//User is a struct  that  represent  a  user  in our  application

type  User struct{

	FullNme string `json:"fullname"`
	UserName string `json:"username"`
	Email  string   `json:"email"`
}

// the Post is a struct  that  represent  a single  post

type Post struct{

	Title string
	Body string
	Author User
}

var data[]Post = []Post{}


func main(){

	router := mux.NewRouter()

	//router.HandleFunc("/test" ,test)
	//router.HandleFunc("/add/{item}",addItems).Methods("GET", "DELETE")
	router.HandleFunc("/posts", getItem).Methods("GET")
	router.HandleFunc("/posts",addItems).Methods("POST")
	router.HandleFunc("/posts/{id}",getPosts).Methods("GET")
	router.HandleFunc("/posts/{id}",updateItem).Methods("PUT")
	router.HandleFunc("/posts/{id}", patchItem).Methods("PATCH")
	router.HandleFunc("/posts/{id}", deleteItem).Methods("DELETE")
	http.ListenAndServe(":9090", router)

}


// func test(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type", "Application/json")
// 	fmt.Println("Response")
// 	json.NewEncoder(w).Encode(struct {
// 		ID string
// 	}{"7"})

// }
func getPosts(w http.ResponseWriter, r * http.Request){
	w.Header().Set("Content-Type", "Application/json")
	//get  the  Id  of the  post  from  the  route  parameter

	 var idParam string =  mux.Vars(r)["id"]

	 id, err := strconv.Atoi(idParam)
	 if err != nil{

	 	w.WriteHeader(400)
	 	w.Write([]byte("ID could  not be  converted  to  integer"))
	 	return
	 }

	 // error checking

	 if id >= len(data){
	 	w.WriteHeader(404)
	 	w.Write([]byte("No data found with  specified ID"))
	 	return
	 }

	 post := data[id]

	 json.NewEncoder(w).Encode(post)
}


func addItems(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "Application/json")

	//routerVariable := mux.Vars(r)["item"]
	//get the Item value from the JSON body

	var newPost Post

	json.NewDecoder(r.Body).Decode(&newPost)

	data = append(data, newPost)

	json.NewEncoder(w).Encode(data)


}

func getItem(w http.ResponseWriter, r * http.Request){

	w.Header().Set("Content-Type", "Application/json")

	fmt.Println("Your details")

	json.NewEncoder(w).Encode(data)

}

func updateItem(w http.ResponseWriter, r * http.Request){
	w.Header().Set("Content-Type", "Application/json")

	var idParam string = mux.Vars(r)["id"]

	id,err := strconv.Atoi(idParam)

	if err != nil{

		w.WriteHeader(400)
		w.Write([]byte("ID could not converted to Integer"))
		return
	}

	//error checking

	if id >= len(data){

		w.WriteHeader(404)
		w.Write([]byte("No data founded with  specified ID"))
		return

	}

	var updatedItem Post

	//updateItem := Post

	json.NewDecoder(r.Body).Decode(&updatedItem)

	data[id] = updatedItem

	json.NewEncoder(w).Encode(updatedItem)

}

func patchItem(w http.ResponseWriter, r * http.Request){

	w.Header().Set("Content-Type", "Application/json")

	var idParam string = mux.Vars(r)["id"]

	id,err := strconv.Atoi(idParam)

	if err != nil{

		w.WriteHeader(400)
		w.Write([]byte("ID could not converted to Integer"))
		return
	}

	//error checking

	if id >= len(data){

		w.WriteHeader(404)
		w.Write([]byte("No data founded with  specified ID"))
		return

	}

	// get the  current  value

	patchdata := data[id]



	json.NewDecoder(r.Body).Decode(&patchdata)


     data[id] = patchdata

	json.NewEncoder(w).Encode(patchdata)

			//another way

	// 	patchdata := &data[id]



	// json.NewDecoder(r.Body).Decode(patchdata)


	// json.NewEncoder(w).Encode(patchdata)
}

func deleteItem(w http.ResponseWriter, r * http.Request){
	w.Header().Set("Content-Type", "Application/json")

	var idParam string = mux.Vars(r)["id"]

	id,err := strconv.Atoi(idParam)

	if err != nil{

		w.WriteHeader(400)
		w.Write([]byte("ID could not converted to Integer"))
		return
	}

	//error checking

	if id >= len(data){

		w.WriteHeader(404)
		w.Write([]byte("No data founded with  specified ID"))
		return

	}

	// Delete the  data  from  the slice 
	//http://github.com/goang/go/wiki/SliceTricks#delete

	data = append(data[:id],data[id+1:]...)

	w.WriteHeader(200)
}
