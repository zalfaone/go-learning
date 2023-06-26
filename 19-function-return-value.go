package main

import "fmt"

func getHello(name string) string {
	return name
}

func functionReturnValue() {
	result := getHello("Leon")
	fmt.Println(result)

	fmt.Println(getHello("Claire"))

}
