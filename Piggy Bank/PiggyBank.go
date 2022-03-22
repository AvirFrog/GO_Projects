package main

import (
	"fmt"
	"math/rand"
)

func main() {

	piggyBank := 0.0

	for piggyBank < 20 {

		switch rand.Intn(3) {

		case 0:
			piggyBank += 0.05

		case 1:
			piggyBank += 0.1

		case 2:

			piggyBank += 0.2

		}

		fmt.Printf("there is $ %5.2f in the piggy bank \n", piggyBank)

	}

}
