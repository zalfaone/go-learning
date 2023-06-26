package main

import "fmt"

func sayHelloTo(firstName string) {
	fmt.Println("Hello ", firstName)
}

func functionParameter() {
	sayHelloTo("Leon")
}

func main() {
	functionParameter()
}
