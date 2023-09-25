package main

import (
	"fmt"
)

func endApp() {
	fmt.Println("Aplikasi Selesai!")
	message := recover()
	fmt.Println("Error dengan message: ", message)
}

func runApp(error bool) {
	// defer akan selalu dieksekusi walaupun ada panic
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(true)
	fmt.Println("Mantap")
}
