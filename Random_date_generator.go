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
		DaysInMonth := 31
		typ := ""
		if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
			typ = "przestepny"
		} else {
			typ = "zwykly"
		}
		switch month {

		case 2:

			if typ == "przestepny" {
				DaysInMonth = 29

			} else {
				DaysInMonth = 28

			}

		case 4, 6, 9, 11:
			DaysInMonth = 30

		}

		day := rand.Intn(DaysInMonth) + 1
		fmt.Printf("Data: %v %v.%v.%v %v \n", era, day, month, year, typ)
	}

}
