package main

import "fmt"

func main() {
	var number, jump int

	fmt.Println("Want me to count how much? ")
	fmt.Scanln(&number)
	fmt.Println("Now enter the jump to Count ")
	fmt.Scanln(&jump)

	if number > 1000 {
		fmt.Println("The number is too large, can crash your Computer!!!")
	} else {
		for counter := 0; counter <= number; counter += jump {
			fmt.Println(counter)
		}
	}

}
