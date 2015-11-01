package main

import "fmt"
import "math/cmplx"

var (
	goIsFun bool = true
	maxInt uint64 = 1<<64 - 1
	complex complex128 = cmplx.Sqrt(-5 + 12i)
)
func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, goIsFun, goIsFun)
	fmt.Printf(f, maxInt, maxInt)
	fmt.Printf(f, complex, complex)

	//type convertion 
	var i int = 8
	var fl float64 = float64(i)
	var u uint = uint(fl)
	fmt.Printf(f, i, i)
	fmt.Printf(f, fl, fl)
	fmt.Printf(f, u, u)
}