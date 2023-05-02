package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Bayu"
	names[1] = "Adi"
	names[2] = "Pratama"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		80,
		90,
		95,
	}

	fmt.Println(values)

	fmt.Println(len(names))
}
