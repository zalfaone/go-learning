package main

import "fmt"

func getHelloChief(name string) string {
	return "Hello Chief " + name
}

func functionValue() {
	hellochief := getHelloChief
	fmt.Println(hellochief("Leon"))
}

func main() {
	functionValue()
}
