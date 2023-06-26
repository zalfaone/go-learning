package main

import "fmt"

func booleanOperator() {
	var ujian = 80
	var absensi = 74

	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi >= 80
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusAbsensi && lulusUjian
	fmt.Println(lulus)

}

func main() {
	booleanOperator()
}
