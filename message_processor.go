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
	}

	replyMessage(e, "Thanks for the image! You can access your image here for a short amount of time: "+imageFilename)
	log.Println("imageId: " + imageFilename)
}
