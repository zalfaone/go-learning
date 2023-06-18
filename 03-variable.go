package main

import "fmt"

func main() {
	var name string

	name = "Leon S Kennedy"
	fmt.Println(name)

	name = "Ada Wong"
	fmt.Println(name)

	var friendname = "Clair Redfield"
	fmt.Println(friendname)

	country := "Indonesia"
	fmt.Println(country)

	country = "United States"
	fmt.Println(country)

	var (
		chapter string = "Chapter 1 - Begin"
		part    int8   = 1
	)
	fmt.Println(chapter)
	fmt.Println(part)
}
