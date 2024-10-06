package main

import "fmt"

func main() {
	var firstName string = "john"

	// var lastName string
	// lastName = "wick"
	lastName := "wick"

	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Printf("halo john wick!\n")
	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println("halo", firstName, lastName+"!")

	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	var fourth, fifth, sixth string = "empat", "lima", "enam"
	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	_ = "belajar Golang"
	_ = "Golang itu mudah"
	name, _ := "john", "wick"

	name := new(string)

	fmt.Println(name)  // 0x20818a220
	fmt.Println(*name) // ""
}
