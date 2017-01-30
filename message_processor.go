package main

import (
	"log"
	"sort"
	"strconv"
	"strings"
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

	//We have to use a string string map because mongo db can only handle strings as keys
	emotionResultMap := make(map[string]string)

	for _, emotionData := range emotionDataSlice {
		emotionResultMap[strconv.Itoa(emotionData.FaceRectangle.Left)] = determineEmotion(emotionData)
	}

	log.Println("Predicted Emotions: ")
	for k, v := range emotionResultMap {
		log.Println("Left Value:", k, "Emotion: ", v)
	}

	updateImage(e.Source.UserId, imageData)
	changeImageUploaded(e.Source.UserId, true)
	updateEmotionData(e.Source.UserId, emotionResultMap)

	// Create a slice to sort the emotion result keys

	var facePositionSlice []int
	for k, _ := range emotionResultMap {
		convertedKey, err := strconv.Atoi(k)
		if err != nil {
			log.Println("Error sorting emotion result keys")
			log.Println(err.Error())
			return
		}
		facePositionSlice = append(facePositionSlice, convertedKey)
	}

	sort.Ints(facePositionSlice)

	pictureDescriptionSlice := []string{"This is a picture of " + imageData.Description.Captions[0].Text + "."}

	firstPersonEmotion := strings.ToLower(emotionResultMap[strconv.Itoa(facePositionSlice[0])])
	numberOfFaces := len(imageData.Faces)

	pictureDescriptionSlice = append(pictureDescriptionSlice, "There appear to be "+strconv.Itoa(numberOfFaces)+" people in this picture.")

	// We will only read emotions for groups of people up to 3.
	if numberOfFaces > 0 && numberOfFaces < 4 {

		switch numberOfFaces {
		case 1:
			firstPersonEmotion := strings.ToLower(emotionResultMap[strconv.Itoa(facePositionSlice[0])])
			if firstPersonEmotion != "" {
				pictureDescriptionSlice = append(pictureDescriptionSlice, "The person in this picture looks "+firstPersonEmotion+".")
			}
		case 2:
			leftPersonEmotion := strings.ToLower(emotionResultMap[strconv.Itoa(facePositionSlice[0])])
			rightPersonEmotion := strings.ToLower(emotionResultMap[strconv.Itoa(facePositionSlice[1])])
			if leftPersonEmotion != "" {
				pictureDescriptionSlice = append(pictureDescriptionSlice, "The person on the left looks "+leftPersonEmotion+".")
			}
			if rightPersonEmotion != "" {
				pictureDescriptionSlice = append(pictureDescriptionSlice, "The person on the right looks "+rightPersonEmotion+".")
			}
		case 3:
			leftPersonEmotion := strings.ToLower(emotionResultMap[strconv.Itoa(facePositionSlice[0])])
			centerPersonEmotion := strings.ToLower(emotionResultMap[strconv.Itoa(facePositionSlice[1])])
			rightPersonEmotion := strings.ToLower(emotionResultMap[strconv.Itoa(facePositionSlice[1])])

			if leftPersonEmotion != "" {
				pictureDescriptionSlice = append(pictureDescriptionSlice, "The person on the left looks "+leftPersonEmotion+".")
			}
			if centerPersonEmotion != "" {
				pictureDescriptionSlice = append(pictureDescriptionSlice, "The person in the middle looks "+centerPersonEmotion+".")
			}
			if rightPersonEmotion != "" {
				pictureDescriptionSlice = append(pictureDescriptionSlice, "The person on the right looks "+rightPersonEmotion+".")
			}
		}

	}
	audioReplyMessage(e, []string{convertToVoice(pictureDescriptionSlice)})
}

func processAudioMessage(e Event) {

	audioFilename, err := downloadAudio(e.Message.Id)
	if err != nil {
		log.Println("Error downloading audio")
		log.Println(err.Error())
		return
	}

	replyMessage(e, []string{"Thanks for the audio file!! You can access your image here for a short amount of time: " + audioFilename})
	log.Println("audioId: " + audioFilename)

}
