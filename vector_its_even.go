package main

import "fmt"

func main() {
	var listSize int

	fmt.Println("--------------------------------")
	fmt.Println("How many number will you write? ")
	fmt.Println("--------------------------------")
	fmt.Scanln(&listSize)

	number := make([]int, listSize)

	for i := 0; i < listSize; i++ {
		fmt.Print("Type one number: ")
		fmt.Scanln(&number[i])

		if number[i]%2 == 0 {
			fmt.Println("This is even!!")
		} else {
			fmt.Println("This number is not even...")
		}

	}

	fmt.Println("-----------------------")
	fmt.Println("You type this numbers: ")
	fmt.Println("-----------------------")
	fmt.Println(number)

}
