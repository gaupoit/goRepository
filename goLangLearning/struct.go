package main 

import (
	"fmt"
	"time"
)

type Bootcamp struct {
	Lat float64
	Lon float64
	Date time.Time
}

type Point struct {
	X, Y int
}

var (
	p1 = Point{1, 2}
	p2 = &Point{1, 2}
	p3 = Point{X: 1}
	p4 = Point{}
)

//accessing fields


func main() {
	myCamp := Bootcamp{
		Lat: 1.2324343,
		Lon: -123.2323,
	}
	myCamp.Date = time.Now()
	fmt.Println(myCamp)	
	fmt.Println(p1, p2, p3, p4)
}
