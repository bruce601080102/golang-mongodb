package api

import (
	"gopkg.in/mgo.v2"
)

type User struct {
	Name  string "bson:`name`"
	Times int    "bson:`times`"
}

type Arg struct {
	Keys   string
	C      *mgo.Collection
	Result User
}
