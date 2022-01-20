package main // every go program must have a main package

import "fmt" // import any other library that is needed after the main package

func main() {
	fmt.Println("Hello World") // Println() to print to the console and move to the next line
	fmt.Print("Hello World\n") // Print() to print to the console and continue on the same line (\n for new line)
	fmt.Printf("Hi, my name is %s. I'm %d years old", "Kayode", 21)
	/*
		Printf() takes two arguments.. the string to output and the values to replace them with.
		 %s for a string,
		 %d for a number
	*/
}

// NOTE: a directory cannot have more than one main() function.
