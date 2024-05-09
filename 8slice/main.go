package main

import "fmt"

func main() {
	fmt.Println("Welcome to slices in golang")

	var fruitsList = []string{"mangoes", "banana", "guava"}
	fmt.Printf("Types of fruitlist is %T\n ", fruitsList)

	fruitsList = append(fruitsList[1:])
	fmt.Println(fruitsList)

}
