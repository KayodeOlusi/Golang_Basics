package main

import "fmt"

func main() {
	// var userEmails map[int]string = make(map[int]string)	// mapping values
	// // use make() for an initial starting point

	// userEmails[1] = "example@gmail.com"
	// userEmails[2] = "example2@gmail.com"

	// fmt.Println(userEmails)

	userEmails := map[int]string{
		1: "example@gmail.com",
		2: "example2@gmail.com",
	}
	userEmails[1] = "example3@gmail.com"

	fmt.Println(userEmails)

	firstEmail := userEmails[1]
	fmt.Println(firstEmail)

	// firstEmail , ok := userEmails[4]		// 	'ok' is a boolean to check if a value exists in a map
	// fmt.Println(firstEmail, ok)		// '', false

	// if _, ok := userEmails[4]; ok { // if userEmail[4] exists
	// 	fmt.Println("Email exists")
	// } else {
	// 	fmt.Println("Email doesn't exist")
	// }

	// delete(userEmails, 2) // delete value with key 2 in userEmails
}
