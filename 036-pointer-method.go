package main

import "fmt"

type Man struct {
	Name string
}

/*
*

	Pada pemanggilan struct Man pada func/method gunakan pointer dengan menambahkan *
	seperti pada func Married
*/
func (man *Man) Married() {
	man.Name = "Mr." + man.Name
	fmt.Println("Di method", man.Name)
}

func main() {
	bayu := Man{"Bayu"}
	bayu.Married()

	fmt.Println(bayu.Name)
}
