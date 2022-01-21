// Data Types

package main

import (
	"fmt"
	"reflect" // reflect is a package for checking data types
)

// Integer  
// int int8 int16 int32 int64
// uint uint8 uint26 uint32 uint64
var age int = 21 // declare the variable and specify the data type


// Float
// float32 float64
var gpa float64 = 4.0

// String
// string
var cake string = "mycake" // use double quotes only for strings

// Boolean
// bool (&& || ! < <= >= !=)
var isTrue bool = age > 21


func main() {
	fmt.Println(reflect.TypeOf(age))
	fmt.Println(reflect.TypeOf(float64(age) * 5.5)) // to perform an operation with two different types, convert them to 													the same type
}