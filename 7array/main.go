package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[2] = "Banana"
	fruitList[3] = "Orange"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var vegList = [5]string{"potato", "tomato", "carrot"}
	fmt.Println(len(vegList))

}
