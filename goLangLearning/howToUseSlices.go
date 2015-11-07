package main

import "fmt"

func main() {
	var s = []int{1,2,3,4,5}
	fmt.Println(s)
	//slices a slice 
	fmt.Println(s[2:4]) // [3, 4]
	fmt.Println(s[:4]) //[1,2,3,4]
	fmt.Println(s[2:])

	//making a slice
	pizzas := make([]string, 3) //make 3 pizzas
	pizzas[0] = "Hawaii"
	pizzas[1] = "Dalat"
	pizzas[2] = "Milk"
	fmt.Printf("%q\n", pizzas)

	//append a slice
	city := []string{}
	vietnam := []string{"Dalat", "Nha Trang"}
	australia := []string{"Sydney", "Melbourne"}
	city = append(city, "New York", "Blue Mountain")
	myCities := append(vietnam, australia...)
	fmt.Printf("%q has lengh is %d\n", myCities, len(myCities))	

	//nil slice 
	var z []int
	fmt.Println(z == nil, len(z), cap(z))
}