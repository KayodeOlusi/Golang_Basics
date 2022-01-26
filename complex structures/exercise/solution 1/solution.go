package main

import "fmt"

func average(num1, num2, num3 float64) (avr float64) {
	avr = num1 + num2 + num3
	return avr
}

// func average(num1, num2, num3 int) float64 {
// 	total := num1 + nu2 + num3
// 	return total / 3
// }

func main() {
	fmt.Println(average(2.3, 2.2, 1.1))
}
