package main

import (
	"fmt"
	"math/rand"
)

func main() {

	guess := rand.Intn(100) + 1
	count := 0
	var myShot int
	var lvl int
	fmt.Print("Hard lvl: ")
	fmt.Scan(&lvl)
	for {
		if lvl == 0 {
			fmt.Println("Hard lvl can't be set to 0")
			fmt.Print("Hard lvl: ")
			fmt.Scan(&lvl)
		} else {
			break
		}
	}

	for {
		fmt.Println("Attempts: ", lvl-count)
		if count < lvl {

			fmt.Print("Your number: ")
			fmt.Scan(&myShot)
			if guess == myShot {
				fmt.Println("You win!")
				break
			} else if myShot == 0 {
				fmt.Println("You must put a number!")

			} else if guess < myShot {
				fmt.Println("To much!")

			} else if guess > myShot {

				fmt.Println("To low!")
			}

		} else {
			fmt.Println("You lose!")
			break
		}

		count++

	}

}
