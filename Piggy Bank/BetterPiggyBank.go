package main

import (
	"fmt"
	"math/rand"
)

func main() {

	piggyBank := 0
	dollars := 0
	cents := 0

	for piggyBank < 2000 {

		switch rand.Intn(3) {

		case 0:
			piggyBank += 5

		case 1:
			piggyBank += 10

		case 2:

			piggyBank += 20

		}

		dollars = piggyBank / 100
		cents = piggyBank % 100

		fmt.Printf("there is $ %d.%02d in the piggy bank \n", dollars, cents)

	}

}
