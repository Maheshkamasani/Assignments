package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}
func main() {
   var Book1 Books    /* Declare Book1 of type Book */
   var Book2 Books    /* Declare Book2 of type Book */
 
   /* book 1 specification */
   Book1.title = "Kloudone"
   Book1.author = "Maheswar Reddy Kamasani"
   Book1.subject = "Golang Tutorials"
   Book1.book_id = 1729

   /* book 2 specification */
   Book2.title = "Assignment"
   Book2.author = "Chinnapa Reddy Kanakanti"
   Book2.subject = "Go Learning"
   Book2.book_id = 1916
 
   /* print Book1 info */
   fmt.Printf( "Book 1 title : %s\n", Book1.title)
   fmt.Printf( "Book 1 author : %s\n", Book1.author)
   fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
   fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)

   /* print Book2 info */
   fmt.Printf( "Book 2 title : %s\n", Book2.title)
   fmt.Printf( "Book 2 author : %s\n", Book2.author)
   fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
   fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id)
}