package main

import "fmt"

func main() {
	var kg float64
	var meterUnit float64

	fmt.Println("How many pounds do you Own? ")
	fmt.Scanln(&kg)

	fmt.Println("How tall are You?")
	fmt.Scanln(&meterUnit)

	fmt.Println(imc_total(kg, meterUnit))

}

func imc_total(kg, meterUnit float64) string {
	imc := kg / (meterUnit * meterUnit)

	if imc >= 18.5 && imc <= 25 {
		return "Congratulations, you are at your ideal Weight :)"

	}

	if imc < 17 {
		return "You are Very Underweight :( "

	}

	if imc >= 17 && imc <= 18.5 {
		return "You're a little Underweight :/ "
	}

	if imc > 25 && imc <= 30 {
		return "You are a little Overweight :/ "

	}

	if imc > 30 && imc <= 35 {
		return "It's in obesity :("

	}

	if imc > 35 && imc <= 40 {
		return "It is in severe Obesity :'("

	}

	if imc > 40 {
		return "Morbid Obesity!!!"

	} else {
		return "..."
	}

}
