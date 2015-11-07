package main

import "fmt"

var pow = []int{1, 2, 4, 8}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	//skip value
	var pow2 = make([]int, 10)
	for i := range pow2 {
		if i % 2 == 0 {
			continue
		}
		pow2[i] = 1 << uint(i)
	}

	//skip index
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}

	//use range for map
	cities := map[string]int{
		"New York": 8336697,
		"Los Angeles": 3857799,
		"Chicago": 2714856,
	}

	for i, v := range cities {
		fmt.Printf("%s has %d inhabitants\n", i, v)
	}
}