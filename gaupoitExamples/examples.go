//this file contains several examples by Golang
package main

import (
	"fmt"
)

func main() {

	createMenu()
	
}

func createMenu() {

	var choice int

	fmt.Printf("Hello, welcome you to GoLang world!!!\n")
	fmt.Printf("We have a lot of functions:\n")
	fmt.Printf("1. Check if number if event or odd\n")
	fmt.Printf("2. Convert degree of F to degree of C\n")
	fmt.Printf("What is your choice: ")
	fmt.Scanf("%d", &choice)

	if (choice == 1) {
		checkNumberIsEvenOrOdd()
	} else if (choice == 2) {
		convertFehrenheitToCelsius()
	}
}

func checkNumberIsEvenOrOdd() {

	fmt.Print("Enter a number: ")
	var number int
	fmt.Scanf("%d", &number)

	if (number % 2 == 0) {
		fmt.Printf("%d is even number\n", number)
	} else {
		fmt.Printf("%d is odd number\n", number)
	}

}

func convertFehrenheitToCelsius() {

	fmt.Print("Enter degree of Fehrenheit: ")
	var fehrenheit float64
	fmt.Scanf("%f", &fehrenheit)

	celcius := ((fehrenheit - 32)*5)/9
	fmt.Printf("%f degree of F equals %f degree of C\n", fehrenheit, celcius)

}