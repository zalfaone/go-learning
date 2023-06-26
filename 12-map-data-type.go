package main

import "fmt"

func mapDataType() {
	person := map[string]string{
		"name":    "John",
		"address": "San Francisco",
	}
	person["email"] = "john@gmail.com"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Harry Potter"
	book["author"] = "JK Rowling"
	book["ups"] = "Salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))

}
