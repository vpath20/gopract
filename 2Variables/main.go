package main

import "fmt"

const LoginToken string = "ekkfkerfjerfvner"

func main() {
	var username string = "Vikrant"
	fmt.Println(username)
	fmt.Printf("Variable is of type:%T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type:%T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type:%T \n", smallVal)

	var smallfloat float64 = 255.443554353545455454545
	fmt.Println(smallfloat)
	fmt.Printf("Variable is of type:%T \n", smallfloat)

	//defaut value and aliases
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("Variable is of type:%T \n", anothervariable)

	//no var style
	numberofuser := 300000
	fmt.Println(numberofuser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type:%T \n", LoginToken)
}
