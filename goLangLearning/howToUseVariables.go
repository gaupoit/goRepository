package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x string = "Hello World"
	fmt.Println(x)

	//another way to declare variable
	y := "Hello Vietnam!"
	fmt.Println(y)
	fmt.Println(reflect.TypeOf(y))

	//another way to define multiple variables
	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	
}