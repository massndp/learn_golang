package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye" + name
}

func main() {
	sayGoodBye := getGoodBye

	result := sayGoodBye("Bayu")
	fmt.Println(result)
	fmt.Println(getGoodBye("Bayoe"))
}
