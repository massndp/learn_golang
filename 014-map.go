package main

import (
	"fmt"
)

func main() {
	person := map[string]string{
		"name":    "Bayu Adi ",
		"address": "Lampung",
	}

	person["title"] = "Programmer" //for adding or changing

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Learn Golang"
	book["author"] = "Bayu"
	book["oops"] = "Salah"
	fmt.Println(book)

	delete(book, "oops")

	fmt.Println(book)
}
