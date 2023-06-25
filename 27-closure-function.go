package main

import "fmt"

func main() {
	name := "Leon"
	counter := 0

	increment := func() {
		// name := "Claire"
		name = "Claire"
		fmt.Println(counter)
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)

}
