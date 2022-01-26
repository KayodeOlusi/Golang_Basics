package main

import "fmt"

func printAge(age int) int {
	fmt.Println(age)
	return age
} // function takes in parameters of a type and the return type must be specified

func printTwo(num int) (int, int) {
	fmt.Println(num)
	return 0, num
} // we can specify how many values we want to return despite a single paramenter

func printSpec(num int) (myAge int, yourAge int) {
	myAge = 15
	yourAge = 16
	return
} // we can give names to the value we want to return and assign a value to them. we return them implicitly with the 		   'return' keyword, which returns the values we assigned

func printAges(ages ...int) {
	for _, value := range ages {
		fmt.Println(value)
	}
} // only specify the return type is the function must return a value
// use the '...' when we don't want to specify how many parameters we pass into the function

func main() {
	printAge(8)               // 8
	printTwo(4)               // 4
	fmt.Println(printSpec(8)) // 15, 16
	printAges(1, 2, 3, 4)
} // main function does not take in any parameter as it only runs our program
