// Life days calculator to this today
package main

import (
	"fmt"
	"time"
)

func main() {
	var day, month, year, totaldays int

	fmt.Println("Write the Day, Month, and Year of birth: ")

	fmt.Print("Day: ")
	fmt.Scanln(&day)

	fmt.Print("Month: ")
	fmt.Scanln(&month)

	fmt.Print("Year: ")
	fmt.Scanln(&year)

	for i := year + 1; i < time.Now().Year(); i++ {
		totaldays += daysyear(i)
	}

	totaldays += dayuntilyear(day, month, year)
	totaldays += dayonyear(time.Now().Day(), int(time.Now().Month()), time.Now().Year())

	fmt.Println("You Have: ", totaldays, " days to life")
}

func daysyear(year int) int {
	if isleap(year) {
		return 366
	} else {
		return 365
	}
}

func dayonyear(day, month, year int) int {
	var totaldays int

	for i := 0; i < month-1; i++ {
		totaldays += daysbymonth(i+1, year)
	}

	return totaldays + day
}

func dayuntilyear(day, month, year int) int {
	var totaldays int

	for i := month + 1; i < 13; i++ {
		totaldays += daysbymonth(i, year)
	}

	totaldays += daysbymonth(month, year) - day

	return totaldays
}

func daysbymonth(month, year int) int {
	if month == 2 && isleap(year) {
		return 29
	}

	var months [12]int

	months[0] = 31
	months[1] = 28
	months[2] = 31
	months[3] = 30
	months[4] = 31
	months[5] = 30
	months[6] = 31
	months[7] = 31
	months[8] = 30
	months[9] = 31
	months[10] = 30
	months[11] = 31

	return months[month-1]
}

func isleap(year int) bool {
	return year%4 == 0
}
