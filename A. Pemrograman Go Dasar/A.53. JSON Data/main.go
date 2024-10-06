package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	var jsonString = `{"Name": "john wick", "Age": 27}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user :", data.FullName)
	fmt.Println("age  :", data.Age)

	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println("user :", data1["Name"])
	fmt.Println("age  :", data1["Age"])

	var data2 interface{}
	json.Unmarshal(jsonData, &data2)

	var decodedData = data2.(map[string]interface{})
	fmt.Println("user :", decodedData["Name"])
	fmt.Println("age  :", decodedData["Age"])

	// 	jsonString = `[
	//     {"Name": "john wick", "Age": 27},
	//     {"Name": "ethan hunt", "Age": 32}
	// ]`

	// 	var dataa []User

	// 	err = json.Unmarshal([]byte(jsonString), &data)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		return
	// 	}

	// 	fmt.Println("user 1:", dataa[0].FullName)
	// 	fmt.Println("user 2:", dataa[1].FullName)

	var object = []User{{"john wick", 27}, {"ethan hunt", 32}}
	jsonData, err = json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	jsonString = string(jsonData)
	fmt.Println(jsonString)
}
