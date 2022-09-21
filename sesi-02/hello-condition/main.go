package main

import "fmt"

func main() {
	var currentYear = 2001
	if age := currentYear - 1998; age < 17 {
		fmt.Printf("kamu belum bile membuat kartu sim => %d", age)
	}
	fmt.Println("\n=======================================")

	var score = 10

	if score > 7 {
		switch {
		case score > 7:
			fmt.Println("perfect")
		default:
			fmt.Println("nice!")
		}
	} else {
		if score == 5 {
			fmt.Println("not bad")
		} else if score == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if score == 0 {
				fmt.Println("Try Harder!")
			}
		}
	}
	fmt.Println("\n=======================================")
}
