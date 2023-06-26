package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func filterSpam(name string) string {
	if name == "Anjing" {
		return "Bola Basket"
	}
	return name
}

func functionAsParameter() {
	sayHelloWithFilter("Leon S Kennedy", filterSpam)
	sayHelloWithFilter("Anjing", filterSpam)

	hellofilter := filterSpam
	sayHelloWithFilter("Claire Redfield", hellofilter)

}

func main() {
	functionAsParameter()
}
