package main

import "fmt"

func main() {
	type noKTP string
	type Married bool

	var ktpBayu noKTP = "132943499192211"
	var marriedStatus Married = true

	fmt.Println(ktpBayu)
	fmt.Println(marriedStatus)
}
