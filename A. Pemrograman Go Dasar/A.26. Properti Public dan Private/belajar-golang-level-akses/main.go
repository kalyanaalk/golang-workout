package main

import (
	"belajar-golang-level-akses/library"
	// . "belajar-golang-level-akses/library"
	"fmt"
)

func main() {
	// library.SayHello("ethan")
	// library.introduce("ethan")

	// var s1 = Student{"ethan", 21}
	// fmt.Println("name ", s1.Name)
	// fmt.Println("grade", s1.Grade)

	// sayHello("ethan")

	fmt.Printf("Name  : %s\n", library.Student.Name)
	fmt.Printf("Grade : %d\n", library.Student.Grade)
}
