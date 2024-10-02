package main

import (
	"fmt"
	"strings"
)

func main() {
	var secret interface{}
	secret = "ethan hunt"
	fmt.Println(secret)
	secret = []string{"apple", "manggo", "banana"}
	fmt.Println(secret)
	secret = 12.4
	fmt.Println(secret)

	// var data map[string]any

	// data = map[string]any{
	// "name": "ethan hunt",
	// "grade": 2,
	// "breakfast": []string{"apple", "manggo", "banana"},
	// }

	secret = 2
	var number = secret.(int) * 10

	fmt.Println(secret, "multiplied by 10 is :", number)
	secret = []string{"apple", "manggo", "banana"}
	var gruits = strings.Join(secret.([]string), ", ")
	fmt.Println(gruits, "is my favorite fruits")

	// type person struct {
	// 	name string
	// 	age int
	//  }
	//  var name = secret.(*person).name
	//  var secret interface{} = &person{name: "wick", age: 27}
	//  fmt.Println(name)

	var person = []map[string]interface{}{
		{"name": "Wick", "age": 23},
		{"name": "Ethan", "age": 23},
		{"name": "Bourne", "age": 22},
	}
	for _, each := range person {
		fmt.Println(each["name"], "age is", each["age"])
	}

	var fruits = []interface{}{
		map[string]interface{}{"name": "strawberry", "total": 10},
		[]string{"manggo", "pineapple", "papaya"},
		"orange",
	}
	for _, each := range fruits {
		fmt.Println(each)
	}
}
