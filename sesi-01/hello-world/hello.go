package main

import "fmt"

func main() {
	/*
		Println()
		Print dan membuat baris baru setelahnya.
	*/
	fmt.Println("Hello")
	fmt.Println("Haiqal", "Ramanizar")
	fmt.Println("======================")

	/*
		Print()
		Print tanpa membuat baris baru.
	*/
	fmt.Print("Haiqal\n")
	fmt.Print("Haiqal", " Ramanizar\n")
	fmt.Print("Haiqal", " ", "Ramanizar\n")
	fmt.Println("======================")

	/*
		Variable with data type
	*/
	var fullname string = "Haiqal"
	var age int = 22
	fmt.Println("Hello ", fullname)
	fmt.Println("Umur kamu", age)
	fmt.Println("======================")

	/*
		Variable without data type
	*/
	var firstname = "Haiqal"
	var gender = "Laki-laki"
	fmt.Printf("%T %T\n", firstname, gender) // %T digunakan untuk mengetahui tipe data variable

	lastname := "Ramanizar"
	birth := "Bogor"
	fmt.Printf("%s %s\n", lastname, birth)
	fmt.Println("======================")

	/*
		Multiple variable declaration
	*/
	var one, two, three string = "1", "2", "3"
	var first, second, third int = 1, 2, 3
	fmt.Printf("%s %s %s\n", one, two, three)
	fmt.Printf("%d %d %d\n", first, second, third)
	fmt.Println("======================")

	/*
		Underscore variable
	*/
	var testUnderscore, testUnderscore2 string
	_, _ = testUnderscore, testUnderscore2 // ketika di-compile maka tidak terjadi compile error

	/*
		Data type
	*/

	// Integer
	var typeOfInt int8 = -12
	var typeOfUint uint8 = 89
	fmt.Printf("tipe %T\n", typeOfInt)
	fmt.Printf("value: %d\n", typeOfInt)
	fmt.Printf("tipe uint %T\n", typeOfUint)
	fmt.Printf("value: %d\n", typeOfUint)

	// Float
	var typeOfFloat32 float32 = 3.63
	fmt.Printf("tipe %T\n", typeOfFloat32)
	fmt.Printf("value: %f\n", typeOfFloat32)
	fmt.Printf("value: %.3f\n", typeOfFloat32)

	// Bool
	var typeOfBool bool = true
	fmt.Printf("tipe %T\n", typeOfBool)
	fmt.Printf("value: %t\n", typeOfBool)

	// String
	var typeOfString string = "test"
	var withBacktick string = `lorem ipsum dolor sit amet
	this is "lorem ipsum"`
	fmt.Printf("tipe %T\n", typeOfString)
	fmt.Printf("value: %s\n", typeOfString)
	fmt.Printf("value with backtic: %s\n", withBacktick)

	/*
		constant
	*/
	const fullName string = "Haiqal Ramanizar"
	fmt.Printf("Hallo %s", fullName)
	fmt.Println("======================")

	/*
		arithmetic operator
	*/
	var value = 2 + 2*3
	fmt.Println(value)

	/*
		operator relational
	*/
	var firstCond bool = 2 > 3
	var secCond bool = "joey" == "Joey"
	var thirdCond bool = 10 != 2.3
	var fourthCond bool = 11 <= 12

	fmt.Println("first condition :", firstCond)
	fmt.Println("second condition :", secCond)
	fmt.Println("third condition :", thirdCond)
	fmt.Println("fourth condition :", fourthCond)
	fmt.Println("======================")

	/*
		logical operator
	*/
	var right = true
	var wrong = false

	var wrongAndRight = wrong && right
	fmt.Printf("wrong && right \t(%T) \n", wrongAndRight)

	var wrongOrRight = wrong || right
	fmt.Printf("wrong || right \t(%T) \n", wrongOrRight)

	var wrongReverse = !wrong
	fmt.Printf("!wrong \t(%T) \n", wrongReverse)
	fmt.Println("======================")
}
