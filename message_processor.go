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
	emotionDataSlice, err := emotionApi(imageFilename)
	if err != nil {
		log.Println("Error calling vision API")
		log.Println(err.Error())
	}

	emotionResultMap := make(map[int]string)

	for _, emotionData := range emotionDataSlice {
		emotionResultMap[emotionData.FaceRectangle.Left] = determineEmotion(emotionData)
	}

	updateImage(e.Source.UserId, imageData)
	changeImageUploaded(e.Source.UserId, true)

	log.Println("Predicted Emotions: ")
	for k, v := range emotionResultMap {
		log.Println("k:", k, "v:", v)
	}

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
