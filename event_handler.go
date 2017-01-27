package main

import (
	"log"
	"strings"
)

func processMessageEvent(e Event) {

	var haveSeenUser bool = false
	var displayName string = ""

	// Get the User ID to check if we've seen this user before
	// Skip this processing if this is not a user
	if e.Source.UserId != "" {

		profile, err := getProfile(e.Source.UserId)
		if err != nil {
			log.Println("ERROR: Could not get profile")
		} else {
			displayName = profile.DisplayName
		}

		if userInDatabase(e.Source.UserId) == false {
			log.Println("User is not database")
			err = addUserToDatabase(e.Source.UserId)
			if err != nil {
				log.Println(err.Error())
				log.Println("Could not add user to database")
			}
		} else {
			log.Println("User is in database")
			haveSeenUser = true
		}
	}

	if e.Message.Type == "text" {
		log.Println(e.Message.Text)

		switch strings.ToLower(e.Message.Text) {
		case "forget me":
			err := removeUserFromDatabase(e.Source.UserId)
			if err != nil {
				replyMessage(e, "Oops, I couldn't forget you!")
				log.Println("Error removing user from database")
			} else {
				replyMessage(e, "Okay, I'll pretend I haven't seen you before!")
			}
		default:
			if haveSeenUser {
				replyMessage(e, "Hello "+displayName+", I've see you before!")
			} else {
				replyMessage(e, "Hello "+displayName+", you're new here, aren't you?")
			}
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
