package main 

import (
	"fmt"
)

func main() {
	//declare a list of variables
	var (
		name string
		age int
		location string
	)
	name = "Thinh"
	age = 10
	location = "VN"
	fmt.Println("declare a list of variables")
	fmt.Printf("Name: %s, Age: %d, Location: %s\n", name, age, location)	

	//declare one by one
	var name2 string
	var age2 int
	var location2 string

	name2 = "Isa"
	age2 = 23
	location2 = "VN"
	fmt.Println("declare one by one")
	fmt.Printf("Name: %s, Age: %d, Location: %s\n", name2, age2, location2)	

	//variable can get the type of initializer 
	var (
		name3 = "Lina"
		age3 = 24
		location3 = "Australia"
	)
	fmt.Println("can get the type of initializer ")
	fmt.Printf("Name: %s, Age: %d, Location: %s\n", name3, age3, location3)	
}