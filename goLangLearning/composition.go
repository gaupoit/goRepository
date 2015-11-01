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
	*User //composition here, change to pointer 
	Point int
}

func newPlayer(id int, name string, location string, point int) *Player {
	return &Player{
		&User{1, name, location},
		point,
	}
}

func main() {
	me := newPlayer(1, "Thinh", "VN", 10)
	fmt.Printf("%+v\n", me)
	fmt.Println(me.Greetings())
}