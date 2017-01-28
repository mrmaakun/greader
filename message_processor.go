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

	log.Println("imageId: " + imageFilename)

	// Flag the user as having sent an image

	imageData, err := visionApi(imageFilename)

	if err != nil {
		log.Println("Error calling vision API")
		log.Println(err.Error())
	}

	updateImage(e.Source.UserId, imageData)
	changeImageUploaded(e.Source.UserId, true)

	replyMessage(e, imageData.Description.Captions[0].Text)

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
