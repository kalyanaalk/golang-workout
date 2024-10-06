package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])

	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	var numbers = [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}

	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	// for i, fruit := range fruits {
	// 	fmt.Printf("nama buah : %s\n", fruit)
	// }

	for _, fruit := range fruits {
		fmt.Printf("nama buah : %s\n", fruit)
	}

	fruits[0] = "apple"
	fruits[1] = "manggo"

	fmt.Println(fruits) // [apple manggo]
}
