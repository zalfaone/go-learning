package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function!")
}

func runApplication(value int) {
	// add defer on top of block code
	defer logging()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result", result)
	// logging()
}

func main() {
	runApplication(0)
}
