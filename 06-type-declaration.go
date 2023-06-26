package main

import "fmt"

func typeDeclaration() {
	type NoKTP string
	type Married bool

	var noKtpLeon NoKTP = "222131212"
	var marriedStatus Married = true
	fmt.Println(noKtpLeon)
	fmt.Println(marriedStatus)
}

func main() {
	typeDeclaration()
}
