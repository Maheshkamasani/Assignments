package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Maheswar Reddy Kamasani ", Desc: "Studying Golang", Content: "Hello Mahesh! Welcome to Kloudone"},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}
func testArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test POST endpoint worked")
}
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}
func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
func main() {
	handleRequests()
}
