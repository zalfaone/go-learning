package main

import "fmt"

func ifExpression() {
	var name = "Ada"

	if name == "Leon" {
		fmt.Println("Hello Leon")
	} else if name == "Claire" {
		fmt.Println("Hello Claire")
	} else {
		fmt.Println("Ini Bukan Leon")
	}

	if len(name) > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Terlalu Panjang")
	}

}

func main() {
	ifExpression()
}
