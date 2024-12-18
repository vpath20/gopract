package main

import "fmt"

func main() {
	fmt.Println("Functions in golang")
	greeter()

	result := adder(3, 5)

	fmt.Println("Result is: ", result)

	proRes, myMessage := proAdder(2, 5, 8, 7)
	fmt.Println("Pro result is: ", proRes)
	fmt.Println("my message is: ", myMessage)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}

	return total, "Hi, Pro Adder"
}

func greeter() {
	fmt.Println("Hi,Hello")

}
