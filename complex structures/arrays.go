package main

import "fmt"

func main() {
	var score [5]float64 // Initialize an array with length of 5 unassigned values
	var nums [3]int

	fmt.Println(score) // [0,0,0,0,0]
	fmt.Println(nums)  // [0,0,0]

	//********************************** Other ways of initializing an array
	// var scored [3]float64 = [3]float64{9, 12.54, 17.6}
	// scores := [3]float64{5.7, 2, 6.09}
	// scoring := [...]float64{6, 3, 7, 0.4, 43.2}
	// scores := [5]float64{1, 2, 3, 4, 5}

	//**********************************Looping through the elements in an array
	// for _, number := range scores {
	// 	fmt.Println(number)
	// }

	// ***************** Using slice and make()

	var myArray [5]int
	var mySlice []int = make([]int, 5, 10)
	// first argument for the type, second for the initial length, third for the 										maximum length
	// make() gives the array a starting length and a maximum amount of elements that can be in the array
	myArray[0] = 1
	mySlice[0] = 1

	fmt.Println(myArray)
	fmt.Println(mySlice)
	fmt.Println(len(myArray))
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	// Slice
	myFood := [5]string{"banana", "soup", "fish", "egg", "rice"}
	otherFood := myFood[1:3]                           // otherFood is an array with values from index 1 to 2 of myFood array
	fruitSlice := append(otherFood, "Cabbage", "Salt") // append() to fill up the remaining space in the sliced array
	// if another value is added, the length is increased and the capacity is doubled

	fmt.Println(otherFood)
	fmt.Println(len(otherFood))        // 2
	fmt.Println(cap(otherFood))        // 4 the capacity is 4 as it still stores the other elements uin myFood array
	fmt.Println(otherFood, fruitSlice) // [soup, fish] [soup, fish, cabbage, salt]
}
