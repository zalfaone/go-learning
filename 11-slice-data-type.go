package main

import "fmt"

func sliceDataType() {
	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	var slice1 = months[4:8]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Mantap")
	fmt.Println(slice3)
	slice3[1] = "Bukan_Desemver"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Leon"
	newSlice[1] = "Kennedy"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	thisArray := [...]int{1, 2, 3, 4}
	thisSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(thisArray)
	fmt.Println(thisSlice)

}

func main() {
	sliceDataType()
}
