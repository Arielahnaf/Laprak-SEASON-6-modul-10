package main

import "fmt"

func main() {
	var b int
	fmt.Print("bilangan: ")
	fmt.Scanln(&b)
	fmt.Printf("Faktor: ")
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Print("\n")
	if b%2 == 0 || b%3 == 0 || b%5 == 0 || b%7 == 0 && b != 1 && b != 2 && b != 3 && b !=
		5 && b != 7 {
		fmt.Println("FALSE")
	} else {
		fmt.Println("TRUE")
	}
}
