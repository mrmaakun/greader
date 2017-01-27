package main

import (
	"log"
)

func processTextMessage() {

}
func processImageMessage(e Event) {
	replyMessage(e, "Thanks for the image! Give me a minute while I check it out.")

	imageFilename, err := downloadImage(e.Message.Id)
	if err != nil {
		log.Println("Error downloading image")
		log.Println(err.Error())
	}

	log.Println("imageId: " + imageFilename)
}
