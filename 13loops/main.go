package main

import "fmt"

func main() {
	fmt.Println("Loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Friday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }

	// for _, day := range days {
	// 	fmt.Printf("Value is %v\n", day)
	// }

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 2 {
			goto lco
		}

		if rogueValue == 5 {
			rogueValue++
			continue
		}

		fmt.Println("Value is ", rogueValue)
		rogueValue++
	}

lco:
	fmt.Println("Jumping at google.com")

}
