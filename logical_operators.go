package main

import "fmt"

func main() {
	var numberA, numberB int

	fmt.Println("Write A then B, and find out if A is greater than B: ")
	fmt.Scanln(&numberA)
	fmt.Scanln(&numberB)

	result := numberA > numberB

	fmt.Println("Its, ", result)

}
