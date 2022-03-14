package main

import (
	"fmt"
	"math/rand"
	"time"
)

var era = "AD"

func main() {

	for count := 1000; count > 0; count-- {
		time.Sleep(time.Second)
		year := rand.Intn(2150) + 1
		month := rand.Intn(12) + 1
		daysInMonth := 31
		typeOfYear := ""
		if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
			typeOfYear = "leap year"
		} else {
			typeOfYear = "year non-leap"
		}
		switch month {

		case 2:

			if typeOfYear == "leap year" {
				daysInMonth = 29

			} else {
				daysInMonth = 28

			}

		case 4, 6, 9, 11:
			daysInMonth = 30

		}

		day := rand.Intn(daysInMonth) + 1
		fmt.Printf("Data: %v %v.%v.%v (Type of year: %v) \n", era, day, month, year, typeOfYear)
	}

}
