package main

import (
	"fmt"
	"strings"
)

func main() {
	// append function
	var fruits = make([]string, 3)

	_ = fruits

	fmt.Printf("%#v", fruits)

	// fruits[0] = "apple"
	// fruits[1] = "banana"
	// fruits[3] = "mango"
	fruits = append(fruits, "apple", "banana", "mango")

	fmt.Printf("%#v", fruits)

	// append with ellipsis (...)
	var fruits1 = []string{"apple", "banana", "mango"}
	var fruits2 = []string{"durian", "pineapple", "starfruit"}

	fruits1 = append(fruits1, fruits2...)

	fmt.Printf("%#v", fruits1)

	// copy elements
	var fruits3 = []string{"apple", "banana", "mango"}
	var fruits4 = []string{"durian", "pineapple", "starfruit"}
	nn := copy(fruits3, fruits4)

	fmt.Println("Fruits1 =>", fruits3)
	fmt.Println("Fruits2 =>", fruits4)
	fmt.Println("Copy elements =>", nn)

	// slicing
	var fruits5 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	var fruits6 = fruits5[1:4]
	fmt.Printf("%#v\n", fruits6)

	var fruits7 = fruits5[0:]
	fmt.Printf("%#v\n", fruits7)

	var fruits8 = fruits5[:3]
	fmt.Printf("%#v\n", fruits8)

	var fruits9 = fruits5[:]
	fmt.Printf("%#v\n", fruits9)

	//slicing & append
	var fruits10 = append(fruits5[:3], "rambutan")
	fmt.Printf("%#v", fruits10)

	//backing array
	var fruits11 = fruits5[2:4]
	fruits11[0] = "rambutan"

	fmt.Println("fruits backing array =>", fruits11)

	//cap array
	fmt.Println("Fruits cap:", cap(fruits5))
	fmt.Println("Fruits len:", len(fruits5))

	fmt.Println(strings.Repeat("#", 30))
	var fruits12 = fruits5[0:3]

	fmt.Println("Fruits cap:", cap(fruits12))
	fmt.Println("Fruits len:", len(fruits12))

	fmt.Println(strings.Repeat("#", 30))
	var fruits13 = fruits5[1:]

	fmt.Println("Fruits cap:", cap(fruits13))
	fmt.Println("Fruits len:", len(fruits13))

	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"
	fmt.Println("cars:", cars)
	fmt.Println("new cars:", newCars)

}
