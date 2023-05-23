package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello Cok!"
	} else {
		return "Hello" + name
	}
}

func main() {
	result := getHello("Bayu Adi P")
	fmt.Println(result)

	fmt.Println(getHello(""))
}
