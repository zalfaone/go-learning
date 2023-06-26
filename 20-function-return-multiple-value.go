package main

import "fmt"

func getFullName() (string, string) {
	return "Leon", "Kennedy"
}

func functionReturnMultipleValue() {
	// firstName, lastName := getFullName()
	// fmt.Println(firstName)
	// fmt.Println(lastName)

	firstName, _ := getFullName()
	fmt.Println(firstName)
}
