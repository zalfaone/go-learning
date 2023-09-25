package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var customer Customer

	customer.Name = "Leon S Kennedy"
	customer.Address = "Racoon City"
	customer.Age = 18

	new_customer := Customer{
		Name:    "Claire",
		Address: "Racoon City",
		Age:     18,
	}

	fmt.Println(customer)
	fmt.Println(new_customer)
}
