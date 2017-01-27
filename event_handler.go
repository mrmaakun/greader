package main

import (
	"log"
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

		if checkIfFirstTimeUser(e.Source.UserId) {
			addUserToDatabase(e.Source.UserId)
		} else {
			haveSeenUser = true
		}
	}

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
