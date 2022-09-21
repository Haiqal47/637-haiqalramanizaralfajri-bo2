package main

import "fmt"

func main() {
	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 1 {
	// 		continue
	// 	}
	// 	if i > 8 {
	// 		break
	// 	}
	// 	fmt.Println("angka", i)
	// }

	// for i := 0; i < 5; i++ {
	// 	for j := i; j < 5; j++ {
	// 		fmt.Print(j, " ")
	// 	}
	// 	fmt.Println()
	// }

outerLoop:
	for i := 0; i < 3; i++ {
		fmt.Println("perulangan ke - ", i+1)
		for j := 0; j < 3; j++ {
			if i == 2 {
				break outerLoop
			}
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}
