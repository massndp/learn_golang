package main

import "fmt"

func main() {
	var name string

	name = "Bayu adi pratama"
	fmt.Println(name)

	name = "bayu"
	fmt.Println(name)

	var friendName = "Dado"
	fmt.Println(friendName)

	// declare variable without key var
	country := "Indonesia"
	fmt.Println(country)

	country = "USA"
	fmt.Println(country)

	// declare multiple variable
	var (
		firstName = "Dado"
		lastName  = "Pridado"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)

}
