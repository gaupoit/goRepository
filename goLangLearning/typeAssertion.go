package main

import (
	"fmt"
	"time"
)	

func timeMap(y interface{}) {
	z, ok := y.(map[string]interface{})
	if ok {
		z["update_time"] = time.Now()
	}
}
func main() {
	foo := map[string]interface{}{
		"gaupo": 25,
	}
	timeMap(foo)
	fmt.Println(foo)
}