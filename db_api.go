package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
)

func userInDatabase(userId string) bool {

	// Connect to Mongo DB
	session, err := mgo.Dial(os.Getenv("MONGO_DB_URL"))
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB(os.Getenv("MONGO_DB_NAME")).C("users")
	result := User{}
	err = c.Find(bson.M{"userid": userId}).One(&result)
	if err != nil {
		log.Println(err.Error())
		return false
	} else {
		return true
	}

}

func addUserToDatabase(userId string) error {

	// Connect to Mongo DB
	session, err := mgo.Dial(os.Getenv("MONGO_DB_URL"))
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB(os.Getenv("MONGO_DB_NAME")).C("users")
	err = c.Insert(&User{userId, false})
	return err

}

func removeUserFromDatabase(userId string) error {

	// Connect to Mongo DB
	session, err := mgo.Dial(os.Getenv("MONGO_DB_URL"))
	if err != nil {
		return err
	}
	defer session.Close()

	c := session.DB(os.Getenv("MONGO_DB_NAME")).C("users")
	err = c.Remove(bson.M{"userid": userId})
	return err
}