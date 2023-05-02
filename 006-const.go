package main

import "fmt"

func main() {
	//const ex 1
	// const firstName string = "Bayu"
	// const lastName string = "Pratama"
	// const value = 1000

	const (
		firstName string = "Bayu"
		lastName  string = "Pratama"
		value            = 1000
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

}
