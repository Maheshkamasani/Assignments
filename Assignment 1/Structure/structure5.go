package main

import "fmt"

func main() {

    u := struct {
        name       string
        occupation string
        age        int
    }{
        name:       "Maheswar Reddy Kamasani",
        occupation: "Trainee @ KLoudone.",
        age:        22,
    }

    fmt.Printf("%s is %d years old and He is a %s\n", u.name, u.age, u.occupation)
}