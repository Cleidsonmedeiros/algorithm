//	Add Numbers
package main

import "fmt"

func main() {
	var numberA, numberB int

	fmt.Println("Write the two numbers to be Summed: ")
	fmt.Scanln(&numberA)
	fmt.Scanln(&numberB)

	fmt.Println("The sum of these numbers is: ", add(numberA, numberB))

}

func add(numberA, numberB int) int {
	return numberA + numberB

}
