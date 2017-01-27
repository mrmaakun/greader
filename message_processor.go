package main

import (
	"log"
)

func processTextMessage() {

}

func processImageMessage(e Event) {

	imageFilename, err := downloadImage(e.Message.Id)
	if err != nil {
		log.Println("Error downloading image")
		log.Println(err.Error())
		return
	}

	replyMessage(e, "Thanks for the image! You can access your image here for a short amount of time: "+imageFilename)
	log.Println("imageId: " + imageFilename)

	// Flag the user as having sent an image
	changeImageUploaded(e.Source.UserId, true)
}

func processAudioMessage(e Event) {

	audioFilename, err := downloadAudio(e.Message.Id)
	if err != nil {
		log.Println("Error downloading audio")
		log.Println(err.Error())
		return
	}

	replyMessage(e, "Thanks for the audio file!! You can access your image here for a short amount of time: "+audioFilename)
	log.Println("audioId: " + audioFilename)

}
