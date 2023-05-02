/**
Type Assertions adalah kemampuan merubah tipe data menjadi tipe data yang diinginkan
Fitur ini sering digunakan ketika bertemu dengan interface kosong
*/

package main

import "fmt"

func random() interface{} {
	// return "Opps"
	return true
}

func main() {
	var result interface{} = random()
	// var resultString int = result.(int)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}
}
