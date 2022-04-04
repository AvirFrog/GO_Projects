package main

import "fmt"

func main() {

	var true_false_yes_no_1_0 string = "zero"
	var bvar bool
	var alert bool = true
	switch true_false_yes_no_1_0 {

	case "true", "yes", "1":
		bvar = true

	case "false", "no", "0":
		bvar = false

	default:
		alert = false

	}

	if alert {
		fmt.Printf("true_false_yes_no_1_0 = %v", bvar)

	} else {

		fmt.Println("true_false_yes_no_1_0 =", true_false_yes_no_1_0)
		fmt.Println("(this variable should contain string:  true/false/yes/no/1/0)")

	}

}
