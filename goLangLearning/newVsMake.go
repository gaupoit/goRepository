package main

import (
	"fmt"
)
//What is the different between: make and new
//
//
//
//1.new : create an instance with zeroes and return the its address
type Point struct {
	X, Y int
}

//2. make: create slides, map, channels only.
var (
	v *[]int = new([]int) //*v = nil
	v2 []int = make([]int, 100) //make a array of 100 ints
)
func main() {
	p := new(Point)
	//or we can write like that
	p2 := &Point{X:0, Y: 0}
	fmt.Println(p)
	fmt.Println(p2)
	fmt.Println(*v)
	fmt.Println(v2)
}


