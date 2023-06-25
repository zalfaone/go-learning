package main

import "fmt"

func getCompleteName() (firstName, lastName string) {
	firstName = "Leon"
	lastName = "Kennedy"
	return
}

func main() {
	firstName, lastName := getCompleteName()
	fmt.Println(getCompleteName())
	fmt.Println(firstName)
	fmt.Println(lastName)
}
