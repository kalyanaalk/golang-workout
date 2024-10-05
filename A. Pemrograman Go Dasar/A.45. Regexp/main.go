package main

import (
	"fmt"
	"regexp"
)

func main() {
	var text = "banana burger soup"
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	var res1 = regex.FindAllString(text, 2)
	fmt.Printf("%#v \n", res1)
	// []string{"banana", "burger"}

	var res2 = regex.FindAllString(text, -1)
	fmt.Printf("%#v \n", res2)
	// []string{"banana", "burger", "soup"}

	var isMatch = regex.MatchString(text)
	fmt.Println(isMatch)
	// true

	var str = regex.FindString(text)
	fmt.Println(str)
	// banana

	var idx = regex.FindStringIndex(text)
	fmt.Println(idx)
	// [0, 6]

	str = text[0:6]
	fmt.Println(str)
	// banana

	var str1 = regex.FindAllString(text, -1)
	fmt.Println(str1)
	// ["banana", "burger", "soup"]

	var str2 = regex.FindAllString(text, 1)
	fmt.Println(str2)
	// ["banana"]

	str = regex.ReplaceAllString(text, "potato")
	fmt.Println(str)
	// "potato potato potato"

	str = regex.ReplaceAllStringFunc(text, func(each string) string {
		if each == "burger" {
			return "potato"
		}
		return each
	})
	fmt.Println(str)
	// "banana potato soup"

	var strr = regex.Split(text, -1)
	fmt.Printf("%#v \n", strr)
	// []string{"", "n", "n", " ", "urger soup"}
}
