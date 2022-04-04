package main

import (
	"fmt"
	"strings"
)

func main() {

	cipher_message := "MC TS OKZHPRGNGB N"
	keyword := "golang"

	cipher_message = strings.ToUpper(strings.Replace(cipher_message, " ", "", -1))
	keyword = strings.ToUpper(strings.Replace(keyword, " ", "", -1))

	message := ""
	index := 0

	for i := 0; i < len(cipher_message); i++ {

		c := cipher_message[i] - 'A'
		k := keyword[index] - 'A'

		c = (c-k+26)%26 + 'A'

		message += string(c)
		index++
		index %= len(keyword)

	}

	fmt.Println(message)

}