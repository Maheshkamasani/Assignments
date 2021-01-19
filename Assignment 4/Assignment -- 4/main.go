package main

import (
	"log"
	"net/http"
	"text/template"
)

//var tmpl = template.Must(template.ParseGlob("templates/*"))

func main() {
	log.Println("Server started on: http://localhost:8080")

	http.HandleFunc("/", Index)
	http.HandleFunc("/index", Index)
	http.HandleFunc("/about", About)
	http.HandleFunc("/contact", Contact)
	http.ListenAndServe(":8080", nil)
}

//index page
func Index(w http.ResponseWriter, res *http.Request) {
	log.Println("run index")
	//index
	temp, err := template.ParseFiles("Index.html")
	if err != nil {

	}
	err = temp.Execute(w, nil)
	if err != nil {

	}
	log.Println("stop index")

}

func About(w http.ResponseWriter, res *http.Request) {
	log.Println("run index")
	//index
	temp, err := template.ParseFiles("about.html")
	if err != nil {

	}
	err = temp.Execute(w, nil)
	if err != nil {

	}
	log.Println("stop index")
}
func Contact(w http.ResponseWriter, res *http.Request) {
	log.Println("run index")
	//index
	temp, err := template.ParseFiles("contact.html")
	if err != nil {

	}
	err = temp.Execute(w, nil)
	if err != nil {

	}
	log.Println("stop index")
}