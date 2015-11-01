package main

import (
	"fmt"
)

//create a Artist struct
type Artist struct {
	Name, Genre string
	Songs int
}

// func newRelease(a Artist) int {
// 	a.Songs++
// 	return a.Songs
// }

func newRelease(a *Artist) int {
	a.Songs++
	return a.Songs	
}

func main() {
	me := &Artist{Name: "Gaupo", Genre: "Pop", Songs: 5}
	fmt.Printf("%s released his %dth song\n", me.Name, newRelease(me))
	fmt.Printf("%s has total of %d songs\n", me.Name, me.Songs)
}

