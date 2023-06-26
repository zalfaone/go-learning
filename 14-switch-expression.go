package main

import "fmt"

func switchExpression() {
	name := "Albert Whesker"

	switch name {
	case "Leon":
		fmt.Println("Hello Leon")
	case "Hunnigan":
		fmt.Println("Hello Hunnigan")
	case "Claire":
		fmt.Println("Hello Claire")
	default:
		fmt.Println("Lu sapa?")
	}

	subjects := "Math"
	switch len(subjects) > 5 {
	case true:
		fmt.Println("Mapel terlalu panjang")
	case false:
		fmt.Println("Mapel pas")
	}

	weapons := [...]string{"M4", "AK47", "MP5", "UZI", "SCAR-L"}
	switch {
	case len(weapons) < 3:
		fmt.Println("Please, bring more weapons bfeora war.")
	case len(weapons) >= 5:
		fmt.Println("Good, you are ready for war.")
	default:
		fmt.Println("You can start war now, or it's better to bring more weapons.")
	}
}

func main() {
	switchExpression()
}
