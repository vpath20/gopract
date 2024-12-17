package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	vikrant := User{"Vikrant", "vikrant@go.org", true, 15}
	fmt.Println(vikrant)
	fmt.Printf("vikrant details are: %+v\n", vikrant)
	fmt.Printf("Name is %v and email is %v.", vikrant.Name, vikrant.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
