package main

import (
	"errors"
	"fmt"
)

func pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi

		return result, nil
	}
}

func main() {
	// var contohError error = errors.New("Opps error")
	// fmt.Println(contohError.Error())

	hasil, err := pembagi(100, 0)

	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {

		fmt.Println("Error", err.Error())
	}

}
