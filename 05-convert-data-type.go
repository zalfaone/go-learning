package main

import "fmt"

func main() {
	var nilai32 int32 = 129
	var nilai64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Leon S Kennedy"
	var e = name[0]
	var eString = string(e)

	fmt.Println(e)
	fmt.Println(eString)
}