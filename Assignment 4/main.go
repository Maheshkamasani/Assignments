package main

import (
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./mahesh")))

	log.Fatal(http.ListenAndServe(":9090", nil))
}