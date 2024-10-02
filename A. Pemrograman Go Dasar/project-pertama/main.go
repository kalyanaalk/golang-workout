package main

import "fmt"

func main() {
	var firstName string = "john"

	lastName := "wick"

	fmt.Printf("halo %s %s!\n", firstName, lastName)

	fmt.Printf("halo john wick!\n")
	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println("halo", firstName, lastName+"!")

	// _ = "belajar Golang"
	// _ = "Golang itu mudah"
	// name, _ := "john", "wick"

	name := new(string)
	fmt.Println(name)  // 0x20818a220
	fmt.Println(*name) // ""

	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644
	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	var message string = "Halo"
	fmt.Printf("message: %s \n", message)

	message = `Nama saya "John Wick".
	Salam kenal.
	Mari belajar "Golang".`

	fmt.Println(message)

	const first string = "john"
	fmt.Print("halo ", first, "!\n")
	const last = "wick"
	fmt.Print("nice to meet you ", last, "!\n")

	var value = (((2+6)%3)*4 - 2) / 3
	var isEqual = (value == 2)
	fmt.Printf("nilai %d (%t) \n", value, isEqual)

	var point = 8
	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}

	var poin = 8840.0
	if percent := poin / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	var poi = 6
	switch poi {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	var po = 6
	switch po {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	var pointy = 6
	switch {
	case pointy == 8:
		fmt.Println("perfect")
	case (pointy < 8) && (pointy > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
		fallthrough
	case point < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	point = 10
	if point > 7 {
		switch point {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point == 5 {
			fmt.Println("not bad")
		} else if point == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point == 0 {
				fmt.Println("try harder!")
			}
		}
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	var i = 0
	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}

	i = 0
	for {
		fmt.Println("Angka", i)
		i++
		if i == 5 {
			break
		}
	}

	var xs = "123" // string
	for i, v := range xs {
		fmt.Println("Index=", i, "Value=", v)
	}
	var ys = [5]int{10, 20, 30, 40, 50} // array
	for _, v := range ys {
		fmt.Println("Value=", v)
	}
	var zs = ys[0:2] // slice
	for _, v := range zs {
		fmt.Println("Value=", v)
	}
	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2} // map
	for k, v := range kvs {
		fmt.Println("Key=", k, "Value=", v)
	}
	// boleh juga baik k dan atau v nya diabaikan
	for range kvs {
		fmt.Println("Done")
	}
	// selain itu, bisa juga dengan cukup menentukan nilai numerik perulangan
	for i := range 5 {
		fmt.Print(i) // 01234
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}
		if i > 8 {
			break
		}
	}

	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}

	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"
	fmt.Println(names[0], names[1], names[2], names[3])

	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	fruits = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}
	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	var numbers = [...]int{2, 3, 2, 4, 3}
	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	var fruit = [4]string{"apple", "grape", "banana", "melon"}
	for i, fruit := range fruit {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	// var fruit = [4]string{"apple", "grape", "banana", "melon"}
	for _, fruit := range fruits {
		fmt.Printf("nama buah : %s\n", fruit)
	}

	// var fruit = [4]string{"apple", "grape", "banana", "melon"}
	for i := 0; i < len(fruit); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}

	var fruity = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruity[0]) // "apple"
	// var fruityA = []string{"apple", "grape"} // slice
	// var fruityB = [2]string{"banana", "melon"}  // array
	// var fruityC = [...]string{"papaya", "grape"}  // array

	// var fruits = []string{"apple", "grape", "banana", "melon"}
	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]
	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]
	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]
	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"
	fmt.Println(fruits)   // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]

	// var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruits)) // 4

	fmt.Println(len(fruits)) // len: 4
	fmt.Println(cap(fruits)) // cap: 4
	aFruits = fruits[0:3]
	fmt.Println(len(aFruits)) // len: 3
	fmt.Println(cap(aFruits)) // cap: 4
	bFruits = fruits[1:4]
	fmt.Println(len(bFruits)) // len: 3
	fmt.Println(cap(bFruits)) // cap: 3

	var fruitss = []string{"apple", "grape", "banana"}
	var cFruitss = append(fruitss, "papaya")

	fmt.Println(fruitss)  // ["apple", "grape", "banana"]
	fmt.Println(cFruitss) // ["apple", "grape", "banana", "papaya"]

	var bFruitss = fruits[0:2]

	fmt.Println(cap(bFruitss)) // 3
	fmt.Println(len(bFruitss)) // 2

	fmt.Println(fruitss)  // ["apple", "grape", "banana"]
	fmt.Println(bFruitss) // ["apple", "grape"]

	fmt.Println(fruitss)  // ["apple", "grape", "papaya"]
	fmt.Println(bFruitss) // ["apple", "grape"]
	fmt.Println(cFruitss) // ["apple", "grape", "papaya"]

	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(n)   // 3

	var fruitts = []string{"apple", "grape", "banana"}
	var aFruitts = fruitts[0:2]
	var bFruitts = fruitts[0:2:2]
	fmt.Println(fruits)
	// ["apple", "grape", "banana"]
	fmt.Println(len(fruitts)) // len: 3
	fmt.Println(cap(fruitts)) // cap: 3
	fmt.Println(aFruitts)
	// ["apple", "grape"]
	fmt.Println(len(aFruitts)) // len: 2
	fmt.Println(cap(aFruitts)) // cap: 3
	fmt.Println(bFruitts)
	fmt.Println(len(bFruitts)) // len: 2
	fmt.Println(cap(bFruitts)) // cap: 2

	// MAP
	var chicken map[string]int
	chicken = map[string]int{}
	chicken["januari"] = 50
	chicken["februari"] = 40
	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei", chicken["mei"])         // mei 0

	// var data map[string]int
	// data["one"] = 1
	// akan muncul error!
	var data = map[string]int{}
	data["one"] = 1
	// tidak ada error

	chicken = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}
	for key, val := range chicken {
		fmt.Println(key, "  \t:", val)
	}

	fmt.Println(len(chicken)) // 4
	fmt.Println(chicken)
	delete(chicken, "januari")
	fmt.Println(len(chicken)) // 3
	fmt.Println(chicken)

	var valuee, isExist = chicken["mei"]
	if isExist {
		fmt.Println(valuee)
	} else {
	}
	fmt.Println("item is not exists")

	var chickens = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}
	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	var dataa = []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}
	for _, dataaa := range dataa {
		fmt.Println(dataaa["gender"], dataaa["name"])
	}
}
