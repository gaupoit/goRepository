package main

import "fmt"

func main() {
	
	const name string = "Hello World"
	fmt.Println(name)

	//now try to modify a const 
	name = "Hello Vietnam" //error : ./howToUseConstant.go:11: cannot assign to name

	fmt.Println(name)
	
}
