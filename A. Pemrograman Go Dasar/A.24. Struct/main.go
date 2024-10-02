package main

import "fmt"

type person struct {
	name string
	age  int
}
type student struct {
	grade int
	person
	age int
}

func main() {
	var s1 = student{}
	s1.name = "wick"
	s1.age = 21
	s1.grade = 2
	fmt.Println("name  :", s1.name)
	fmt.Println("age   :", s1.age)
	fmt.Println("age   :", s1.person.age)
	fmt.Println("grade :", s1.grade)

	s1.person.age = 22
	fmt.Println(s1.person.age)

	var p1 = person{name: "wick", age: 21}
	s1 = student{person: p1, grade: 2}
	fmt.Println("name  :", s1.name)
	fmt.Println("age   :", s1.age)
	fmt.Println("grade :", s1.grade)
	var s11 = struct {
		person
		grade int
	}{}
	s11.person = person{"wick", 21}
	s11.grade = 2

	fmt.Println("name  :", s11.person.name)
	fmt.Println("age   :", s11.person.age)
	fmt.Println("grade :", s11.grade)

	var allStudents = []person{
		{name: "Wick", age: 23},
		{name: "Ethan", age: 23},
		{name: "Bourne", age: 22},
	}
	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}

	var student struct {
		person
		grade int
	}
	student.person = person{"wick", 21}
	student.grade = 2

}
