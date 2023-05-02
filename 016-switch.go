package main

import "fmt"

func main() {
	name := "Bayu"

	switch name {
	case "Bayu":
		fmt.Println("Hello Bayu")

	case "Wahyu":
		fmt.Println("Hello Wahyu")

	default:
		fmt.Println("Hello guys")
	}

	//switch with short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	//switch without condition
	form := len(name)
	switch {
	case form > 10:
		fmt.Println("Hei, terlalu panjang!")
	case form > 5:
		fmt.Println("Hei, masih panjang lho")
	default:
		fmt.Println("Nah ini benar")

	}
}
