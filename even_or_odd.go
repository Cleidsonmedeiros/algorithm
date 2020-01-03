package main

import "fmt"

func main() {
	var number int

	fmt.Println("Write a number and I'll tell you if it's Even or Odd: ")
	fmt.Scanln(&number)
	fmt.Println(evenOdd(number))

}

func evenOdd(number int) string {
	if number%2 == 0 {
		return "Its Even!!"
	} else {
		return "Its Odd"

	}

}
