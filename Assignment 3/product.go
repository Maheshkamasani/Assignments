package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Name  string `json:"Name"`
	Price string `json:"Price"`
	Size  string `json:"Size"`
}

type Products []Product

func allProducts(w http.ResponseWriter, r *http.Request) {
	products := Products{
		Product{Name: "Samsung", Price: "25000", Size: "16.5"},
	}
	fmt.Println("Endpoint Hit: All Products Endpoint")
	json.NewEncoder(w).Encode(products)

}
func testProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " Test POST endpoint Hit")

}
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page Endpont Hit")

}
func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/products", allProducts).Methods("GET")

	log.Fatal(http.ListenAndServe(":8082", myRouter))
}
func main() {
	handleRequest()
}