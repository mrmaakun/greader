package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
)

func getUserFromDatabase(userId string) (User, error) {

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
		return result, err
	} else {
		return result, err
	}

}

func addUserToDatabase(userId string) (User, error) {

	addUser := User{}

	// Connect to Mongo DB
	session, err := mgo.Dial(os.Getenv("MONGO_DB_URL"))
	if err != nil {
		return addUser, err
	}

	addUser = User{userId, false, ImageInformation{}, map[string]string{}}
	defer session.Close()
	c := session.DB(os.Getenv("MONGO_DB_NAME")).C("users")
	err = c.Insert(addUser)
	return addUser, err

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

func changeImageUploaded(userId string, imageUploaded bool) error {

	// Connect to Mongo DB
	session, err := mgo.Dial(os.Getenv("MONGO_DB_URL"))
	if err != nil {
		return err
	}
	defer session.Close()

	c := session.DB(os.Getenv("MONGO_DB_NAME")).C("users")
	err = c.Update(bson.M{"userid": userId}, bson.M{"$set": bson.M{"imageuploaded": imageUploaded}})
	return err

}

func updateImage(userId string, imageData ImageInformation) error {

	// Connect to Mongo DB
	session, err := mgo.Dial(os.Getenv("MONGO_DB_URL"))
	if err != nil {
		return err
	}
	defer session.Close()

	c := session.DB(os.Getenv("MONGO_DB_NAME")).C("users")
	err = c.Update(bson.M{"userid": userId}, bson.M{"$set": bson.M{"imagedata": imageData}})
	return err

}

func updateEmotionData(userId string, emotionData map[string]string) error {

	// Connect to Mongo DB
	session, err := mgo.Dial(os.Getenv("MONGO_DB_URL"))
	if err != nil {
		return err
	}
	defer session.Close()

	c := session.DB(os.Getenv("MONGO_DB_NAME")).C("users")
	err = c.Update(bson.M{"userid": userId}, bson.M{"$set": bson.M{"emotiondata": emotionData}})
	return err

}
