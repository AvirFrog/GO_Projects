package main

import (
	"fmt"
	"strings"
)

func main() {

	message := "Go is betterthan C"
	keyword := "golang"

	message = strings.ToUpper(strings.Replace(message, " ", "", -1))
	keyword = strings.ToUpper(strings.Replace(keyword, " ", "", -1))

	index := 0
	cipher_message := ""

	for i := 0; i < len(message); i++ {

		c := message[i]

		if c >= 'A' && c <= 'Z' {

			c -= 'A'

			k := keyword[index] - 'A'

			c = (c+k)%26 + 'A'

			index++

			index %= len(keyword)

		}

		cipher_message += string(c)
	}

	fmt.Println(cipher_message)
}
