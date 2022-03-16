package main

import (
	"fmt"
	"math/rand"
)

func main() {

	date := "Date"
	space_line := "Space line"
	days := "Days"
	type1 := "Type"
	price := "Price"
	speed := "Speed"

	const second_per_day = 86400
	const distance = 62100000

	fmt.Printf("%-15v %-20v %-5v %-20v $(M) %-10v %v km/s\n\n", date, space_line, days, type1, price, speed)
	fmt.Printf("==========================================================================================\n\n")
	date = "13.10.2020"

	for count := 0; count < 50; count++ {

		switch rand.Intn(10) {

		case 0:

			space_line = "JanuszSpace"

		case 1:

			space_line = "SpaceX"

		case 2:

			space_line = "NASA"

		case 3:

			space_line = "PolSA"

		case 4:

			space_line = "PolSpace"

		case 5:

			space_line = "DarekRocketPol"

		case 6:

			space_line = "Rocket Science"

		case 7:

			space_line = "Rakieta PL"

		case 8:

			space_line = "PolMoon"

		case 9:

			space_line = "RakietPol"

		}

		if rand.Intn(2) == 1 {

			type1 = "in two directions"

		} else {

			type1 = "one way"

		}

		speed1 := rand.Intn(15) + 16
		duration := distance / speed1 / second_per_day
		price1 := 20.0 + speed1

		fmt.Printf("%-15v %-20v %-5v %-20v $(M) %-10v %v km/s\n\n", date, space_line, duration, type1, price1, speed1)
	}

}
