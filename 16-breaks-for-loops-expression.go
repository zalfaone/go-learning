package main

import "fmt"

func breakForLoopExpression() {
	for i := 0; i <= 10; i++ {
		if i == 7 {
			break
		}
		fmt.Println("Perulangan ke", i)
	}

	weapons := map[string]string{
		"meele":      "Knive",
		"mid-range":  "Shotgun",
		"long-range": "Sniper",
	}

	for i, value := range weapons {
		if i == "mid-range" {
			break
		}
		fmt.Println("Index", i, "=", value)
	}

	for i, value := range weapons {
		if i == "mid-range" {
			continue
		}
		fmt.Println("Index", i, "=", value)
	}

}
