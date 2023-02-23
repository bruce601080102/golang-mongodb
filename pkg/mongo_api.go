package api

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func ConnectionDB() *mgo.Database {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	db := session.DB("admin") //root user is created in the admin authentication database and given the role of root.
	return db
}

func ReadDB(argDB *Arg) int {
	err := argDB.C.Find(bson.M{"name": argDB.Keys}).One(&argDB.Result)
	if err != nil {
		return 1
	}
	return 0
}

func InsertDB(argDB Arg) error {
	err := argDB.C.Insert(&User{Name: argDB.Keys, Times: argDB.Result.Times})

	return err
}

func Update(times int, argDB Arg) {
	data := bson.M{"$set": bson.M{"times": times}}
	selector := bson.M{"name": argDB.Result.Name}
	_ = argDB.C.Update(selector, data)
}

func Remove(argDB Arg) {
	argDB.C.Remove(bson.M{"name": argDB.Keys})
}
