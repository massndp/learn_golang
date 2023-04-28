package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)

	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" || name == "Asu" {
		return "***"
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Bayoe", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
	sayHelloWithFilter("Asu", spamFilter)
}
