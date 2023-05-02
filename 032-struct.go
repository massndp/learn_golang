/**
~ Struct digunakan untuk menggabungkan nol a/ lebih tipe data dalam satu kesatuan
~ Struct disimpan dalam field
~ sederhananya struct adalah kumpulan dari filed
*/

package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// func sayHi(customer Customer, name string) {
// 	fmt.Println("Hello", name, "My name is", customer.Name)
// }

// cara pemanggilan struct method
func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	// cara penulisan struct pertama
	var bayu Customer
	bayu.Name = "Bayu Adi"
	bayu.Address = "Lampung"
	bayu.Age = 27

	// cara pemanggilan func sayHi 1
	// sayHi(bayu, "Dino")

	// cara pemanggilan func sayHi2 dengan struct method
	bayu.sayHi("Dino")

	// fmt.Println(bayu)
	// fmt.Println(bayu.Name)
	// fmt.Println(bayu.Address)
	// fmt.Println(bayu.Age)

	// cara penulisan struct kedua.
	// adi := Customer{
	// 	Name:    "Adi",
	// 	Address: "Jakarta",
	// 	Age:     27,
	// }
	// fmt.Println(adi)

	// cara penulisan struct ketiga
	// *) pada cara ini harus diurutkan berdasarkan urutan data pada structnya
	// elon := Customer{"Elon", "Bandung", 30}
	// fmt.Println(elon)
}
