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
	fmt.Printf("3. Convert feet into meters\n")
	fmt.Printf("What is your choice: ")
	fmt.Scanf("%d", &choice)

	if (choice == 1) {
		checkNumberIsEvenOrOdd()
	} else if (choice == 2) {
		convertFehrenheitInToCelsius()
	} else if (choice == 3) {
		CovertFeetInToMeter()
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

func convertFehrenheitInToCelsius() {

	fmt.Print("Enter degree of Fehrenheit: ")
	var fehrenheit float64
	fmt.Scanf("%f", &fehrenheit)

	celcius := ((fehrenheit - 32)*5)/9
	fmt.Printf("%f degree of F equals %f degree of C\n", fehrenheit, celcius)

}

func CovertFeetInToMeter() {
	
	fmt.Printf("Enter the feet: ")
	var feet float64
	fmt.Scanf("%f", &feet)

	meters := feet*0.3048
	fmt.Printf("%f feet equals to %f meters\n", feet, meters)

}