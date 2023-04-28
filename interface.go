/**
~Interface adalah tipe data abstract, dan tidak memiliki implementasi langsung
~sebuah interface berisikan definisi-definisi method
~Biasanya interface digunakan sebagai kontrak
*/

package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var bayu Person
	bayu.Name = "Bayu"

	sayHello(bayu)

	cat := Animal{
		Name: "Kucing",
	}

	sayHello(cat)
}
