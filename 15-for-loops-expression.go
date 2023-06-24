package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	boss := []string{"T Virus", "C Virus", "G Virus", "T Veronica Virus", "Las Plaga", "Evelyn Virus"}

	for i := 0; i < len(boss); i++ {
		fmt.Println(boss[i])
	}

	for _, value := range boss {
		// fmt.Println("Index", i, "=", value)
		fmt.Println(value)
	}

	for i, value := range boss {
		fmt.Println("Index", i, "=", value)
		// fmt.Println(value)
	}
}
