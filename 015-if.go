package main

import "fmt"

func main() {
	name := "Bayu Adi P"

	if name == "Bayu" {
		fmt.Println("Hello Bayu Adi")
	} else if name == "Wahyu" {
		fmt.Println("Hello Wahyu")
	} else if name == "Andra" {
		fmt.Println("Hello Andra")
	} else {
		fmt.Println("Hi!")
	}

	//if short statment

	//normal ex
	// var lenght = len(name)
	// if lenght > 5 {
	// 	fmt.Println("Terlalu Panjang")
	// } else {
	// 	fmt.Println("Nama sudah benar")
	// }

	//short statment ex
	if lenght := len(name); lenght < 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
