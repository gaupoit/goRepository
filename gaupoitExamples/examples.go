//this file contains several examples by Golang
package main

import (
	"fmt"
)

func main() {

	checkNumberIsEvenOrOdd()
	
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
