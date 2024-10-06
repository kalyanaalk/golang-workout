package main

import "fmt"

func main() {
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	var mmessage string = "Halo"
	fmt.Printf("message: %s \n", mmessage)

	mmessage = `Nama saya "John Wick".
Salam kenal.
Mari belajar "Golang".`

	fmt.Println(mmessage)

}
