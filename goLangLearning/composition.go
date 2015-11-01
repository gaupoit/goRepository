package main

import "fmt"

type User struct {
	Id int
	Name string
	Location string
}

func (u *User) Greetings() string{
	return fmt.Sprintf("Hello. I am %s from %s", u.Name, u.Location)
}

type Player struct {
	User //composition here
	Point int
}

func main() {
	me := Player{}
	me.Id = 1
	me.Name = "Thinh"
	me.Location = "VN"
	me.Point = 0
	fmt.Printf("%+v", me)
	fmt.Println(me.Greetings())
}