package main

import "fmt"

func constant() {
	const (
		firstname string = "Leon"
		lastname         = "Kennedy"
		value            = 1000
	)

	fmt.Println(firstname)
}
