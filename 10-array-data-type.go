package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Leon"
	names[1] = "Scaffold"
	names[2] = "Kennedy"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var products = [3]string{
		"Toothpaste",
		"Bookshelf",
		"Playstation",
	}
	fmt.Println(products)
	fmt.Println(len(products))

	products[0] = "Spoon"
	fmt.Println(products)
}
