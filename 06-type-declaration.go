package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpLeon NoKTP = "222131212"
	var marriedStatus Married = true
	fmt.Println(noKtpLeon)
	fmt.Println(marriedStatus)
}
