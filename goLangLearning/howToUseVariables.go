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

}