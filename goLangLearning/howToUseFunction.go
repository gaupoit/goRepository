package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func location(city string) (string, string) {
	var region string
	var continent string

	switch city {
		case "Los Angeles", "LA", "Santa Monica":
			region, continent = "California", "North Ameria"
		case "New York", "NYC":
			region, continent = "New Yor", "North Ameria"
		default:
			region, continent = "Unknow", "Unknow"
	}
	return region, continent
}

func main() {
	fmt.Println(add(42, 43))
	region, continent := location("New York")
	fmt.Printf("Matt lives in %s, %s\n", region, continent)
}