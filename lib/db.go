package lib 

import (
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func ConnectDb(uri string) *mgo.Session {
	session, err := mgo.Dial(uri)
	if err != nil {
		panic(err)
	}
	return session
}