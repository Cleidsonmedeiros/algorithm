package main

import "fmt"

func main() {
	var numberA, numberB int

	fmt.Print("Type the first number: ")
	fmt.Scanln(&numberA)

	fmt.Print("Type the second number: ")
	fmt.Scanln(&numberB)

	if numberA == 0 || numberB == 0 {
		fmt.Println("Invalid input!")
		return
	}

	fmt.Println("The MDC between", numberA, "and", numberB, "is:", mdc(numberA, numberB))
}

func mdc(numberA, numberB int) int {
	restoAtual := numberA % numberB

	for restoAtual != 0 {
		restoAtual = numberA % numberB
		numberA = numberB

		if restoAtual == 0 {
			return numberB
		}

		numberB = restoAtual
	}

	return numberB
}
