package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Counter ke ", counter)
		counter++
	}

	//for with statement
	for count := 1; count <= 10; count++ {
		fmt.Println("Perulangan ke", count)
	}

	//loop with slice
	slice := []string{"Bayu", "Adi", "Pratama"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//for with range
	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	//looping with range, when variable not use. Change with _ (underscore)
	for _, val := range slice {
		fmt.Println(val)
	}

}
