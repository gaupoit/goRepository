package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type fakeString struct {
	content string
}

//function implement Stringer
func (s *fakeString) String() string {
	return s.content
}

func printString(value interface{}) {
	switch str := value.(type) {
		case string:
			fmt.Println(str)
		case Stringer:
			fmt.Println(str.String())
		default:
			fmt.Println("Unknow")	
	}
}

func main() {
	s := &fakeString{"Hello world!"}
	printString(s)
	printString("Hello string")
}