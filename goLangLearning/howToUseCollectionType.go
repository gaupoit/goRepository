package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	//another way to create a array
	var a2 = [2]string{"Hello", "World"}
	fmt.Println(a2[0], a2[1])
	fmt.Println(a2)
	fmt.Printf("%q\n", a2)

	//use ellipsie to use a implicit length
	a3 := [...]string{"Hello", "World", "!"}
	fmt.Printf("%q\n", a3)

	//multi-dimension
	var di [3][2]string
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			di[i][j] = fmt.Sprintf("row %d -- column %d", i, j)
		}
	} 

	fmt.Printf("%q\n", di)
}