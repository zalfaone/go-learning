package main

import "fmt"

func sumNumbers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	total := sumNumbers(10, 10, 20, 50)
	fmt.Println(total)

	slice := []int{10, 20, 30}
	total = sumNumbers(slice...)
	fmt.Println(total)
}
