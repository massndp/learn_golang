/**
~ Secara deafult di go semua variable di passing by value bukan by refrence
~ Artinya, jika ingin mengirim sebuah variable kedalam function, method atau variable lain,
sebenarnya yang dikirim adalah duplikasi value nya

~ Pointer adalah kemampuan membuat refrence ke lokasi data di memory yang sama,
tanpa menduplikasi data yang sudah ada
~ Sederhananya, dengan kemampuan pointer, kita bisa membuat by refrence
*/

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

/*
*

	Penerapan pointer pada function
*/
func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var address1 Address = Address{"Tanggamus", "Lampung", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1
	var address4 = new(Address)

	address2.City = "Bandar Lampung"

	*address2 = Address{"Palembang", "Sumatera Selatan", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address4)

	/**
	Penerapan pointer pada function
	*/
	var alamat = Address{
		City:     "Metro",
		Province: "Lampung",
		Country:  "",
	}

	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}
