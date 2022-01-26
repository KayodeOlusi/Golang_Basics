package main

import (
	"fmt"
	"reflect"
)

func average(nums ...float64) float64 {
	total := 0.0
	for _, number := range nums {
		total += number
	}
	return total / float64(len(nums))
}

func main() {
	fmt.Println(average(2, 6, 9))
	fmt.Println(reflect.TypeOf(average()))
}
