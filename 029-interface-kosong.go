package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return 2
	} else {
		return "Ups"
	}
}

func main() {
	// cara menampilkan interface kosong
	var data interface{} = Ups(2)
	fmt.Println(data)
}
