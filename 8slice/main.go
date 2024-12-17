package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in golang")

	var fruitsList = []string{"mangoes", "banana", "guava"}
	fmt.Printf("Types of fruitlist is %T\n ", fruitsList)

	fruitsList = append(fruitsList[:3])
	fmt.Println(fruitsList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	highScores = append(highScores, 624, 320, 123)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)

	//how to remove a value from slices based on index

	var subjects = []string{"math", "english", "biology", "physics"}
	fmt.Println(subjects)

	var index int = 2
	subjects = append(subjects[:index], subjects[index+1:]...)
	fmt.Println(subjects)
}
