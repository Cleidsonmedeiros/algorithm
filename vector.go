package main

import "fmt"

func main() {
	var number [6]int
	var sum int

	for i := 0; i < 6; i++ {

		fmt.Println(number)
		fmt.Print("Fill in the table with Numbers: ")
		fmt.Scanln(&number[i])

		sum = sum + number[i]

	}

	fmt.Println("The sum of these numbers is: ", sum)

}
