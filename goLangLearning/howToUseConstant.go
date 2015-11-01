package main

import "fmt"

func main() {
	
	const name string = "Hello World"
	const PI = 3.14
	const (
		StatusOK = 200
		StatusCreated = 201
		StatusAccepted = 202
	)
	fmt.Println(name)

	//now try to modify a const 
	name = "Hello Vietnam" //error : ./howToUseConstant.go:11: cannot assign to name

	fmt.Println(name)
	
}
