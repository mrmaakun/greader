package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
)

func processMessageEvent(e Event) {

	var haveSeenUser bool = false
	var displayName string = ""

	// Connect to Mongo DB
	session, err := mgo.Dial(os.Getenv("MONGO_DB_URL"))
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB(os.Getenv("MONGO_DB_NAME")).C("users")

	// Get the User ID to check if we've seen this user before
	if e.Source.UserId != "" {
		// Get the profile to store in the db
		profile, err := getProfile(e.Source.UserId)
		if err != nil {
			log.Println("Get profile failed. Putting blank profile data into the DB")
		} else {
			// Record Name
			displayName = profile.DisplayName
		}

		result := User{}
		err = c.Find(bson.M{"userid": e.Source.UserId}).One(&result)
		if err != nil {

			// The user is not found in the database so we will add them.
			log.Println(err.Error())

			err = c.Insert(&User{e.Source.UserId, false, profile})
			if err != nil {
				log.Fatal(err)
			}
		} else {
			haveSeenUser = true
		}
	}

	/*

		DisplayName   string `json:"displayName,omitempty"`
		UserId        string `json:"userId,omitempty"`
		PictureUrl    string `json:"pictureUrl,omitempty"`
		StatusMessage string `json:"statusMessage,omitempty"`
	*/

	if e.Message.Type == "text" {
		log.Println(e.Message.Text)

		if haveSeenUser {
			replyMessage(e, "Hello "+displayName+", I've see you before!")
		} else {
			replyMessage(e, "Hello "+displayName+", you're new here, aren't you?")
		}
	}
}

func handleEvent(e Event) {

	switch e.Type {

	case "message":
		processMessageEvent(e)
	default:
		log.Println("Invalid Event Type")
		return
	}

}
