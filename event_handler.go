package main

import (
	"log"
	"strings"
)

func processMessageEvent(e Event) {

	var haveSeenUser bool = false
	var displayName string = ""
	var currentUserData User = User{}

	// Get the User ID to check if we've seen this user before
	// Skip this processing if this is not a user
	if e.Source.UserId != "" {

		profile, err := getProfile(e.Source.UserId)
		if err != nil {
			log.Println("ERROR: Could not get profile")
		} else {
			displayName = profile.DisplayName
		}

		currentUserData, err = getUserFromDatabase(e.Source.UserId)
		if err != nil {
			log.Println("User is not database")
			currentUserData, err = addUserToDatabase(e.Source.UserId)
			if err != nil {
				log.Println(err.Error())
				log.Println("Could not add user to database")
			}
		} else {
			log.Println("User is in database")
			haveSeenUser = true
		}
	}

	switch e.Message.Type {
	case "text":
		log.Println(e.Message.Text)

		switch strings.ToLower(e.Message.Text) {
		case "forget me":
			err := removeUserFromDatabase(e.Source.UserId)
			if err != nil {
				replyMessage(e, []string{"Oops, I couldn't forget you!"})
				log.Println("Error removing user from database")
			} else {
				replyMessage(e, []string{"Okay, I'll pretend I haven't seen you before!"})
			}
		default:
			if currentUserData.ImageUploaded == true {

				switch strings.ToLower(e.Message.Text) {
				case "no":
					replyMessage(e, []string{"Okay, let's forget about this image!"})
					changeImageUploaded(e.Source.UserId, false)
				case "yes":
					replyMessage(e, []string{currentUserData.ImageData.Description.Captions[0].Text})
				default:
					replyMessage(e, []string{"It looks like you sent me an image. Do you want to know anything about it?"})

				}
			} else {
				if haveSeenUser {
					replyMessage(e, []string{"Hello " + displayName + ", I've see you before! Why don't you send me an image so that I can tell you about it?"})
				} else {
					replyMessage(e, []string{"Hello " + displayName + ", I've never seen you around before. Nice to meet you! If you send me an image, I will tell you about it!"})
				}
			}
		}
	case "image":
		processImageMessage(e)
	case "audio":
		processAudioMessage(e)
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
