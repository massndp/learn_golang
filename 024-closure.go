package main

import "fmt"

func main() {
	name := "Bayu"
	counter := 0

	increment := func() {
		// pendeklarasian seperti ini akan menggantikan var name yang ada di atas func increment
		name = "Budi"
		// maka jika tidak ingin terganti harus di deklarasikan variable baru
		name := "Adi"
		counter++
		fmt.Println("Increment")
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)

}
