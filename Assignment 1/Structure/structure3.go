package main

import "fmt"

type person struct {
    name string
    age  int
}

func newPerson(name string) *person {

    p := person{name: name}
    p.age = 22
    return &p
}

func main() {

    fmt.Println(person{"Mahesh", 22})

    fmt.Println(person{name:"Divya", age: 96})

    fmt.Println(person{name: "Vishnu", age: 18})

    fmt.Println(&person{name: "Chinna", age: 22})

    fmt.Println(newPerson("Bala"))

    s := person{name: "Madhan", age: 21}
    fmt.Println(s.name)

    sp := s
    fmt.Println(sp.age)
}