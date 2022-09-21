package main

import (
	"fmt"
	"strings"
)

func main() {
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	var testStringArray = [3]string{"arif", "andi", "adam"}
	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", testStringArray)

	var fruits = [3]string{"apel", "pisang", "mangga"}
	for i, v := range fruits {
		fmt.Printf("index: %d, value: %s\n", i, v)
	}
	fmt.Println("==================================")

	fmt.Println(strings.Repeat("#", 25))
	fmt.Println("==================================")

	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	for _, arr := range balances {
		fmt.Printf("%d \n", arr)
		for _, v := range arr {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
	fmt.Println("==================================")

}
