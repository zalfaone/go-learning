package main

import "fmt"

func sayHello() {
	fmt.Println("Hello Leon!")
}

func function() {
	for i := 0; i < 10; i++ {
		sayHello()
	}
}
