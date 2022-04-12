package main

import "fmt"

func kelvin2celsius(k float64) float64 {
	k -= 273.15

	return k
}

func kelvin2fahrenheit(k float64) float64 {
	k = (k * 1.8) - 459.67

	return k
}

func kelvin2rankine(k float64) float64 {
	k *= 1.8

	return k
}

func celsius2kelvin(c float64) float64 {
	c += 273.15

	return c
}

func celsius2fahrenheit(c float64) float64 {
	c = (c * 1.8) + 32

	return c
}

func celsius2rankine(c float64) float64 {
	c = (c + 273.15) * 1.8

	return c
}

func fahrenheit2kelvin(f float64) float64 {
	f = ((f - 32.0) / 1.8) + 273.15

	return f

}

func fahrenheit2celsius(f float64) float64 {
	f = (f - 32.0) / 1.8

	return f
}

func fahrenheit2rankine(f float64) float64 {
	f += 459.67

	return f
}

func rankine2kelvin(r float64) float64 {
	r /= 1.8

	return r
}

func rankine2celsius(r float64) float64 {
	r = (r - 491.67) / 1.8

	return r
}

func rankine2fahrenheit(r float64) float64 {
	r -= 459.67

	return r
}

func main() {

	test := 36.6

	fmt.Printf("%v Kelvin degrees is %v Celsius degrees \n", test, kelvin2celsius(test))
	fmt.Printf("%v Kelvin degrees is %v Fahrenheit degrees \n", test, kelvin2fahrenheit(test))
	fmt.Printf("%v Kelvin degrees is %v Rankine degrees \n", test, kelvin2rankine(test))
	fmt.Println("\n")
	fmt.Printf("%v Celsius degrees is %v Kelvin degrees \n", test, celsius2kelvin(test))
	fmt.Printf("%v Celsius degrees is %v Fahrenheit degrees \n", test, celsius2fahrenheit(test))
	fmt.Printf("%v Celsius degrees is %v Rankine degrees \n", test, celsius2rankine(test))
	fmt.Println("\n")
	fmt.Printf("%v Fahrenheit degrees is %v Kelvin degrees \n", test, fahrenheit2kelvin(test))
	fmt.Printf("%v Fahrenheit degrees is %v Celsius degrees \n", test, fahrenheit2celsius(test))
	fmt.Printf("%v Fahrenheit degrees is %v Rankine degrees \n", test, fahrenheit2rankine(test))
	fmt.Println("\n")
	fmt.Printf("%v Rankine degrees is %v Kelvin degrees \n", test, rankine2kelvin(test))
	fmt.Printf("%v Rankine degrees is %v Celsius degrees \n", test, rankine2celsius(test))
	fmt.Printf("%v Rankine degrees is %v Fahrenheit degrees \n", test, rankine2fahrenheit(test))

}
