package main

import "fmt"

func getHelloChief(name string) string {
	return "Hello Chief " + name
}

func main() {
	hellochief := getHelloChief
	fmt.Println(hellochief("Leon"))
}
