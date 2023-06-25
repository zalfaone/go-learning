package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello ", firstName)
}

func main() {
	sayHelloTo("Leon")
}
