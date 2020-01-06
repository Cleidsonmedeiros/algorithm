package main

import "fmt"

func main() {
	var number, jump int

	fmt.Println("Want me to count how much? ")
	fmt.Scanln(&number)
	fmt.Println("Now enter the jump to Count ")
	fmt.Scanln(&jump) // jumping number "important"

	if number > 1000 {
		fmt.Println("This number is too large, can crash your computer")
	} else {
		fmt.Println("From 0 to", number, "Jumping from", jump, "Numbers, Are these numbers:")
		counter(number, jump)

	}
}

func counter(number, jump int) {
	for i := 0; i <= number; i += jump {
		fmt.Println(i)
	}

}
