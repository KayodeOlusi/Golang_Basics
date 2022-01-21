package main

import "fmt"

func main() {
	sentence := "It is Morning"	// declare the variable

	for index, letter := range sentence {	// loop through the variable and check if the index of the letter is odd
		if index % 2 != 0 {	
			fmt.Println(string(letter))		// use string() to convert the number-type string to string-type
		}
	}
}

// once a value is declared and not used, an error is thown. Use "_" to ignore. 